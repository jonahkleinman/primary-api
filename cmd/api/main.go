package main

import (
	"primary-api/pkg/config"
	"primary-api/pkg/database"
	"primary-api/pkg/database/models"
)

func main() {
	cfg := config.New()

	database.Connect(cfg.Database)
	models.AutoMigrate()
}
