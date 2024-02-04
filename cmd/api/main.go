package main

import (
	"github.com/VATUSA/primary-api/pkg/config"
	go_chi "github.com/VATUSA/primary-api/pkg/go-chi"
	"net/http"
)

func main() {
	cfg := config.New()

	//database.DB = database.Connect(cfg.Database)
	//models.AutoMigrate(database.DB)

	r := go_chi.New(cfg)

	go_chi.Testers(r)

	http.ListenAndServe(":8080", r)
}
