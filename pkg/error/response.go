package error

import (
	"citypair/pkg/formatter"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Err            error                       `json:"-"`               // low-level runtime error
	HTTPStatusCode int                         `json:"-"`               // http response status code
	Message        string                      `json:"error,omitempty"` // application-level error message, for debugging
	ValidationErr  []formatter.ValidationError `json:"validation_error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func InvalidParameter(err error, msg string) *ErrResponse {
	if msg == "" {
		msg = "Invalid parameter"
	}
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        msg,
	}
}

func BadRequest(err error, msg string) *ErrResponse {
	if msg == "" {
		msg = "bad request"
	}

	response := &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Message:        msg,
	}

	if v, ok := err.(validator.ValidationErrors); ok {
		response.ValidationErr = formatter.DescribeValidationError(v)
	}

	return response
}
