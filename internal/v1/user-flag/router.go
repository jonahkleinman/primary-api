package user_flag

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListUserFlag)
	r.Post("/", CreateUserFlag)
	r.Route("/{UserFlagID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetUserFlag)
		r.Put("/", UpdateUserFlag)
		r.Delete("/", DeleteUserFlag)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "UserFlagID")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		UserFlagID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		userFlag := &models.UserFlag{ID: uint(UserFlagID)}
		if err = userFlag.Get(database.DB); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "userFlag", userFlag)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFlagCtx(r *http.Request) *models.UserFlag {
	return r.Context().Value("userFlag").(*models.UserFlag)
}
