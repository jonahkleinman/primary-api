package disciplinary_log_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VATUSA/primary-api/internal/v1/disciplinary_log"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/disciplinary", func(r chi.Router) {
		r.Post("/", disciplinary_log.CreateDisciplinaryLogEntry)
		r.Get("/", disciplinary_log.ListDisciplinaryLog)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(disciplinary_log.GetDisciplinaryLogCtx)
			r.Get("/", disciplinary_log.GetDisciplinaryLog)
			r.Put("/", disciplinary_log.UpdateDisciplinaryLog)
			r.Patch("/", disciplinary_log.PatchDisciplinaryLog)
			r.Delete("/", disciplinary_log.DeleteDisciplinaryLog)
		})
	})
	return r
}

func TestDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	t.Run("create", func(t *testing.T) {
		tests := []struct {
			name     string
			body     string
			expected int
		}{
			{"valid request", `{"cid":123456,"entry":"Test entry","vatusa_only":false}`, http.StatusCreated},
			{"invalid cid", `{"cid":0,"entry":"Test entry","vatusa_only":false}`, http.StatusBadRequest},
			{"empty entry", `{"cid":123456,"entry":"","vatusa_only":false}`, http.StatusBadRequest},
			{"missing field", `{"cid":123456,"vatusa_only":false}`, http.StatusBadRequest},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest("POST", "/disciplinary", strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/json")

				rr := httptest.NewRecorder()
				r.ServeHTTP(rr, req)

				defer rr.Body.Close()
				assert.Equal(t, tt.expected, rr.Code)
			})
		}
	})

	t.Run("list", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/disciplinary", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		defer rr.Body.Close()
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("get", func(t *testing.T) {
		dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
		database.DB.Create(&dle)

		req := httptest.NewRequest("GET", "/disciplinary/"+dle.ID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		defer rr.Body.Close()
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("update", func(t *testing.T) {
		dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
		database.DB.Create(&dle)

		req := httptest.NewRequest("PUT", "/disciplinary/"+dle.ID, strings.NewReader(`{"cid":654321,"entry":"Updated entry","vatusa_only":true}`))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		defer rr.Body.Close()
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("patch", func(t *testing.T) {
		dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
		database.DB.Create(&dle)

		req := httptest.NewRequest("PATCH", "/disciplinary/"+dle.ID, strings.NewReader(`{"entry":"Patched entry"}`))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		defer rr.Body.Close()
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("delete", func(t *testing.T) {
		dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
		database.DB.Create(&dle)

		req := httptest.NewRequest("DELETE", "/disciplinary/"+dle.ID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		defer rr.Body.Close()
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
}
