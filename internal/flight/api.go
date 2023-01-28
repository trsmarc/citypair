package flight

import (
	"net/http"

	"citypair/pkg/log"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func RegisterHandlers(r chi.Router, logger log.Logger, validate *validator.Validate) {
	r.Mount("/", RegisterHTTPHandlers(NewFlightHTTP(logger, validate)))
}

func RegisterHTTPHandlers(http HTTP) http.Handler {
	r := chi.NewRouter()
	r.Post("/calculate", http.GetCityPair)
	return r
}
