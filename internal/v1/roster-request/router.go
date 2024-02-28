package roster_request

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListRosterRequest)
	r.Post("/", CreateRosterRequest)
	r.Route("/{RosterRequestID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetRosterRequest)
		r.Put("/", UpdateRosterRequest)
		r.Delete("/", DeleteRosterRequest)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "RosterRequestID")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		RosterRequestID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		rosterRequest := &models.RosterRequest{ID: uint(RosterRequestID)}
		if err = rosterRequest.Get(); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "rosterRequest", rosterRequest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRosterRequestCtx(r *http.Request) *models.RosterRequest {
	return r.Context().Value("rosterRequest").(*models.RosterRequest)
}
