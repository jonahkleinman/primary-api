package models

import (
	"github.com/VATUSA/primary-api/pkg/database"
	"log"
)

func AutoMigrate() {
	err := database.DB.AutoMigrate(
		&Facility{},
		&User{},
		&ActionLogEntry{},
		&DisciplinaryLogEntry{},
		&Document{},
		&FacilityLogEntry{},
		&FAQ{},
		&Feedback{},
		&News{},
		&Notifications{},
		&RatingChange{},
		&Roster{},
		&RosterRequest{},
		&Flag{},
		&UserRole{},
		&Role{},
	)
	if err != nil {
		log.Fatal("[Database] Migration Error:", err)
	}
}
