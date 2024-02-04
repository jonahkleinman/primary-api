package user

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database"
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
	//r.Get("/search", SearchUsers)

	r.Route("/{CID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetUser)
		r.Put("/", UpdateUser)
		r.Delete("/", DeleteUser)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := chi.URLParam(r, "CID")
		// TODO - if "" then check for x-user header
		// TODO - Check Authentication here
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
		err = user.Get(database.DB)
		if err != nil {
			render.Render(w, r, utils.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
