package go_chi

import (
	"github.com/VATUSA/primary-api/pkg/config"
	middleware2 "github.com/VATUSA/primary-api/pkg/middleware"
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

	Testers(r)

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

	// Use NotGuest middleware on get route
	r.Get("/guest", func(w http.ResponseWriter, r *http.Request) {
		middleware2.NotGuest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Passed"))
		})).ServeHTTP(w, r)
	})
}
