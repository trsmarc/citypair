package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func RegisterHandlers(r chi.Router) {
	r.Get("/health", Check)
}

func Check(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "passed")
}
