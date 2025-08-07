package engine

import (
	"context"
	"os"
	"time"

	"src/engine/server"
)

type Option func(o *options)

type options struct {
	ctx  context.Context
	sigs []os.Signal

	stopTimeout time.Duration
	servers     []server.Server
}

func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

func WithStopTimeout(t time.Duration) Option {
	return func(o *options) { o.stopTimeout = t }
}

func WithServer(srv ...server.Server) Option {
	return func(o *options) { o.servers = srv }
}
