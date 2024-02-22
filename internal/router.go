package internal

import (
	v1 "github.com/VATUSA/primary-api/internal/v1"
	"github.com/VATUSA/primary-api/pkg/config"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/VATUSA/primary-api/internal/docs"
)

// @title           VATUSA Internal
// @version         0.1
// @description     VATUSAs internal API for use by other internal VATUSA services.
// @termsOfService  http://swagger.io/terms/

// @contact.name   VATUSA Support
// @contact.url    http://www.swagger.io/support
// @contact.email  vatusa6@vatusa.net

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /internal/v1

func Router(r chi.Router, cfg *config.Config) {
	r.Route("/internal", func(r chi.Router) {
		v1.Router(r, cfg)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/internal/swagger/doc.json"),
	))
}
