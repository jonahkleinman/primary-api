package faq

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListFAQ)
	r.Post("/", CreateFAQ)

	r.Route("/{FAQID}", func(r chi.Router) {
		r.Use(Ctx)
		r.Get("/", GetFAQ)
		r.Put("/", UpdateFAQ)
		r.Patch("/", PatchFAQ)
		r.Delete("/", DeleteFAQ)
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "FAQID")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		faqID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		faq := &models.FAQ{ID: uint(faqID)}
		if err := faq.Get(database.DB); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "faq", faq)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetFAQCtx(r *http.Request) *models.FAQ {
	return r.Context().Value("faq").(*models.FAQ)
}
