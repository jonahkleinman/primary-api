package user

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListUsers)
	r.Post("/", CreateUser)

	r.Route("/{CID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetUser)
		r.Put("/", UpdateUser)
		r.Patch("/", PatchUser)
		r.Delete("/", DeleteUser)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := chi.URLParam(r, "CID")
		if cid == "" {
			render.Render(w, r, utils.ErrNotFound)
			return
		}

		CID, err := strconv.ParseUint(cid, 10, 64)
		if err != nil {
			render.Render(w, r, utils.ErrNotFound)
			return
		}

		user := &models.User{CID: uint(CID)}
		err = user.Get()
		if err != nil {
			render.Render(w, r, utils.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserCtx(r *http.Request) *models.User {
	return r.Context().Value("user").(*models.User)
}
