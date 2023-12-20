package models

import (
	"time"
)

type Roster struct {
	ID         uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID        uint      `json:"cid" example:"1293257"`
	Facility   string    `json:"facility" example:"ZDV"`
	OIs        string    `json:"operating_initials" example:"RP"`
	Home       bool      `json:"home" example:"true"`
	Visiting   bool      `json:"visiting" example:"false"`
	Status     string    `json:"status" example:"Active"` // Active, LOA
	Mentor     bool      `json:"mentor" example:"false"`
	Instructor bool      `json:"instructor" example:"false"`
	CreatedAt  time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt  time.Time `json:"deleted_at" example:"2021-01-01T00:00:00Z"` // Soft Deletes for logging
}
