package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrConflict(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 409,
		StatusText:     "Duplicate Key",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var (
	ErrNotFound        = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
	ErrBadRequest      = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad request"}
	ErrInternalServer  = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}
	ErrInvalidFacility = &ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid facility"}
	ErrInvalidRole     = &ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid role"}
	ErrInvalidCID      = &ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid CID"}
)
