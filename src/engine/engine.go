package engine

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Engine struct {
	opts options
	ctx  context.Context
	cancel context.CancelFunc
}

func New(opts ...Option) *Engine {
	o := options{
		ctx:  context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)
	return &Engine{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

func (e *Engine) Run() error {
	eg, egCtx := errgroup.WithContext(e.ctx)
	wg := sync.WaitGroup{}

	for _, srv := range e.opts.servers {
		server := srv
		eg.Go(func() error {
			<-egCtx.Done()
			stopCtx := e.opts.ctx
			if e.opts.stopTimeout > 0 {
				var stopCancel context.CancelFunc
				stopCtx, stopCancel = context.WithTimeout(stopCtx, e.opts.stopTimeout)
				defer stopCancel()
			}
			return server.Stop(stopCtx)
		})

		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return server.Start(e.opts.ctx)
		})
	}
	wg.Wait()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, e.opts.sigs...)
	eg.Go(func() error {
		select {
		case <-egCtx.Done():
			return nil
		case <-quit:
			return e.Stop()
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (e *Engine) Stop() error {
	if e.cancel != nil {
		e.cancel()
	}
	return nil
}
