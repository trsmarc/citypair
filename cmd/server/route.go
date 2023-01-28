package server

import (
	"citypair/internal/flight"
	"citypair/internal/healthcheck"
	"citypair/pkg/log"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Routing(logger log.Logger) chi.Router {
	validate = validator.New()

	r := chi.NewRouter()

	healthcheck.RegisterHandlers(r)

	r.Route("/v1", func(r chi.Router) {
		flight.RegisterHandlers(r, logger, validate)
	})

	return r
}
