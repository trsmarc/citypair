package flight

import (
	"context"
	"errors"

	"citypair/pkg/log"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetCityPair(ctx context.Context, request *GetCityPairRequest) (string, string, error)
}

type flightService struct {
	validator *validator.Validate
	logger    log.Logger
}

func NewService(validate *validator.Validate, logger log.Logger) Service {
	return flightService{validator: validate, logger: logger}
}

func (s flightService) GetCityPair(ctx context.Context, req *GetCityPairRequest) (string, string, error) {
	var src, dest string

	if err := req.Validate(s.validator); err != nil {
		s.logger.Errorf("Request validation failed on get city pair : %s", err.Error())
		return "", "", err
	}

	pairs := make(map[string]string, len(req.Flights))
	departures := make(map[string]bool, len(req.Flights))
	arrivals := make(map[string]bool, len(req.Flights))

	for _, pair := range req.Flights {
		if len(pair) != 2 {
			return "", "", errors.New("invalid pair input")
		}
		pairs[pair[0]] = pair[1]
		departures[pair[0]] = true
		arrivals[pair[1]] = true
	}

	for k, _ := range departures { //nolint:gosimple
		if !arrivals[k] {
			src = k
			break
		}
	}

	if src == "" {
		return "", "", errors.New("circulate flight")
	}

	sorted := make([]string, 0)
	next := src
	for {
		sorted = append(sorted, next)
		next = pairs[next]
		if next == "" {
			break
		}
	}

	if len(sorted)-1 != len(pairs) {
		return "", "", errors.New("flight is not connected")
	}

	dest = sorted[len(sorted)-1]

	return src, dest, nil
}

type GetCityPairRequest struct {
	Flights [][]string
}

func (request GetCityPairRequest) Validate(validator *validator.Validate) error {
	return validator.Struct(request)
}
