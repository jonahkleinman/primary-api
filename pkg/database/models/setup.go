package models

import (
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Facility{},
		&User{},
		&ActionLogEntry{},
		&DisciplinaryLogEntry{},
		&Document{},
		&FacilityLogEntry{},
		&FAQ{},
		&Feedback{},
		&News{},
		&Notification{},
		&RatingChange{},
		&Roster{},
		&RosterRequest{},
		&UserFlag{},
		&UserRole{},
		&Role{},
	)
	if err != nil {
		log.Fatal("[Database] Migration Error:", err)
	}
}
