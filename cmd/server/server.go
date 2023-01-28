package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"citypair/internal/config"
	"citypair/pkg/log"

	"github.com/pkg/errors"
)

func Start(cfg *config.Config, r http.Handler, logger log.Logger) error {
	serverErrors := make(chan error, 1)

	server := http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	go func() {
		logger.Infof("http server listening on :%s", cfg.Server.Port)
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		logger.Infof("start shutdown: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			logger.Infof("graceful shutdown did not complete in %v : %v", 10, err)
			err = server.Close()
			return err
		}

		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
