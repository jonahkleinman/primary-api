package document

import (
	"context"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Router(r chi.Router) {
	r.Get("/", ListDocuments)
	r.Post("/", CreateDocument)

	r.Route("/{Facility}", func(r chi.Router) {
		r.Get("/", ListDocumentsByFac)
		r.Route("/{Category}", func(r chi.Router) {
			r.Get("/", ListDocumentsByFacByCat)
			r.Route("/{DocumentID}", func(r chi.Router) {
				r.Use(Ctx)
				r.Get("/", GetDocument)
				r.Put("/", UpdateDocument)
				r.Patch("/", PatchDocument)
				r.Delete("/", DeleteDocument)
			})
		})
	})
}

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "DocumentID")
		if id == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		DocumentID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		document := &models.Document{ID: uint(DocumentID)}
		if err = document.Get(database.DB); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "document", document)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDocumentCtx(r *http.Request) *models.Document {
	return r.Context().Value("document").(*models.Document)
}
