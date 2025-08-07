package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	basesrv "src/engine/server"
	"src/internal/config"
	"src/internal/server/http/router"
)

type server struct {
	httpsrv *http.Server
}

func NewServer(cfg config.HTTP) basesrv.Server {
	srv := &server{}

	srv.httpsrv = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Setup(),
	}

	return srv
}

func (srv *server) Start(ctx context.Context) error {
	if err := srv.httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (srv *server) Stop(ctx context.Context) error {
	err := srv.httpsrv.Shutdown(ctx)
	if err != nil {
		if ctx.Err() != nil {
			err = srv.httpsrv.Close()
		}
	}
	return err
}
