package main

import (
	"github.com/VATUSA/primary-api/pkg/config"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"log"
)

func main() {
	cfg := config.New()

	database.Connect(cfg.Database)
	models.AutoMigrate()

	// Create all the possible user roles
	roles := []models.Role{
		{
			Role: "ATM",
			Name: "Air Traffic Manager",
		},
		{
			Role: "DATM",
			Name: "Deputy Air Traffic Manager",
		},
		{
			Role: "TA",
			Name: "Training Administrator",
		},
		{
			Role: "WM",
			Name: "Webmaster",
		},
		{
			Role: "AWM",
			Name: "Web Team Member",
		},
		{
			Role: "EC",
			Name: "Events Coordinator",
		},
		{
			Role: "AEC",
			Name: "Events Team Member",
		},
		{
			Role: "FE",
			Name: "Facility Engineer",
		},
		{
			Role: "AFE",
			Name: "Facilities Team Member",
		},
		{
			Role: "TMU",
			Name: "Traffic Management Operator",
		},
		{
			Role: "USA1",
			Name: "Division Director",
		},
		{
			Role: "USA2",
			Name: "Deputy Director Air Traffic Services",
		},
		{
			Role: "USA3",
			Name: "Deputy Director Training Services",
		},
		{
			Role: "USA4",
			Name: "Deputy Director Support Services",
		},
		{
			Role: "USA5",
			Name: "Events Manager",
		},
		{
			Role: "USA6",
			Name: "Technical Manager",
		},
		{
			Role: "USA7",
			Name: "Staff Development Manager",
		},
		{
			Role: "USA8",
			Name: "Training Services Manager",
		},
		{
			Role: "USA9",
			Name: "Training Content and Curriculum Manager",
		},
	}

	for _, role := range roles {
		if err := role.Create(); err != nil {
			log.Fatalf("error creating role '%s' error: %v", role.Name, err)
		}
	}
}
