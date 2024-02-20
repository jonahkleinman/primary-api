package internal

import (
	v1 "github.com/VATUSA/primary-api/internal/v1"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Route("/internal", func(r chi.Router) {
		v1.Router(r)
	})
}
