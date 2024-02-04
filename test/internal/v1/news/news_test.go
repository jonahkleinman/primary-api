package news_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VATUSA/primary-api/internal/v1/news"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/news", func(r chi.Router) {
		r.Post("/", news.CreateNews)
		r.Get("/", news.ListNews)
		r.Route("/{newsID}", func(r chi.Router) {
			r.Use(news.Ctx)
			r.Get("/", news.GetNews)
			r.Put("/", news.UpdateNews)
			r.Delete("/", news.DeleteNews)
		})
	})
	return r
}

func TestCreateNews(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("POST", "/news", strings.NewReader(`{"title":"Test News","content":"This is a test news"}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestListNews(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/news", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetNews(t *testing.T) {
	r := setupRouter()

	news := &models.News{ID: 1, Title: "Test News", Description: "This is a test news"}
	database.DB.Create(&news)

	req, _ := http.NewRequest("GET", "/news/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateNews(t *testing.T) {
	r := setupRouter()

	news := &models.News{ID: 1, Title: "Test News", Description: "This is a test news"}
	database.DB.Create(&news)

	req, _ := http.NewRequest("PUT", "/news/1", strings.NewReader(`{"title":"Updated Test News","content":"This is an updated test news"}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteNews(t *testing.T) {
	r := setupRouter()

	news := &models.News{ID: 1, Title: "Test News", Description: "This is a test news"}
	database.DB.Create(&news)

	req, _ := http.NewRequest("DELETE", "/news/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
