package go_chi

import (
	"github.com/VATUSA/primary-api/internal/v1/news"
	"github.com/VATUSA/primary-api/internal/v1/user"
	"github.com/VATUSA/primary-api/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
)

func New(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(cors.Handler(NewCors(cfg)))

	r.Route("/internal/v1", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			user.Router(r)
		})

		r.Route("/news", func(r chi.Router) {
			news.Router(r)
		})
	})

	return r
}

func Testers(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong."))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
}
