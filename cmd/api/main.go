package main

import (
	"github.com/VATUSA/primary-api/pkg/config"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
)

func main() {
	cfg := config.New()

	db := database.Connect(cfg.Database)
	models.AutoMigrate(db)
}
