package models

import (
	"log"
	"primary-api/pkg/database"
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
