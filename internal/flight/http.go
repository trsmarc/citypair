package flight

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	e "citypair/pkg/error"
	"citypair/pkg/log"

	"github.com/go-chi/render"
)

type HTTP interface {
	GetCityPair(w http.ResponseWriter, r *http.Request)
}

type flightHTTP struct {
	srv    Service
	Logger log.Logger
}

func NewFlightHTTP(logger log.Logger, validator *validator.Validate) HTTP {
	return flightHTTP{srv: NewService(validator, logger), Logger: logger}
}

func GetFlightHTTP(srv Service, logger log.Logger) HTTP {
	return flightHTTP{srv: srv, Logger: logger}
}

func getCityPairRequestDecoder(r *http.Request) (*GetCityPairRequest, error) {
	var req *GetCityPairRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func (h flightHTTP) GetCityPair(w http.ResponseWriter, r *http.Request) {
	request, err := getCityPairRequestDecoder(r)
	if err != nil {
		h.Logger.With(r.Context()).Errorf("request parsing error : %s", err)
		if err = render.Render(w, r, e.BadRequest(err, "bad request")); err != nil {
			h.Logger.With(r.Context()).Errorf("unable to render payload : %s", err)
			return
		}
		return
	}

	src, dest, err := h.srv.GetCityPair(r.Context(), request)
	if err != nil {
		h.Logger.With(r.Context()).Errorf("get city pair error : %s", err)
		if err = render.Render(w, r, e.BadRequest(err, err.Error())); err != nil {
			h.Logger.With(r.Context()).Errorf("unable to render payload : %s", err)
			return
		}
		return
	}

	render.JSON(w, r, &GetCityPairResponse{Result: []string{src, dest}})
}
