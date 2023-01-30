package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func RegisterHandlers(r chi.Router) {
	r.Get("/health", Check)
}

func Check(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, &HealthCheck{Status: "Passed"})
}
