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

func TestCreateDisciplinaryLogEntry(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("POST", "/disciplinary", strings.NewReader(`{"cid":123456,"entry":"Test entry","vatusa_only":false}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestListDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/disciplinary", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
	database.DB.Create(&dle)

	req, _ := http.NewRequest("GET", "/disciplinary/"+dle.ID, nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
	database.DB.Create(&dle)

	req, _ := http.NewRequest("PUT", "/disciplinary/"+dle.ID, strings.NewReader(`{"cid":654321,"entry":"Updated entry","vatusa_only":true}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPatchDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
	database.DB.Create(&dle)

	req, _ := http.NewRequest("PATCH", "/disciplinary/"+dle.ID, strings.NewReader(`{"entry":"Patched entry"}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteDisciplinaryLog(t *testing.T) {
	r := setupRouter()

	dle := &models.DisciplinaryLogEntry{CID: 123456, Entry: "Test entry", VATUSAOnly: false}
	database.DB.Create(&dle)

	req, _ := http.NewRequest("DELETE", "/disciplinary/"+dle.ID, nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
