package roster

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListRoster)
	r.Post("/", CreateRoster)
	r.Route("/{RosterID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetRoster)
		r.Put("/", UpdateRoster)
		r.Delete("/", DeleteRoster)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "RosterID")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		RosterID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		roster := &models.Roster{ID: uint(RosterID)}
		if err = roster.Get(); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "roster", roster)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRosterCtx(r *http.Request) *models.Roster {
	return r.Context().Value("roster").(*models.Roster)
}
