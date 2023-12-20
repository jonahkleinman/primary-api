package models

import (
	"time"
)

type FAQ struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Facility  string    `json:"facility" example:"ZDV"`
	Question  string    `json:"question" example:"Why shouldn't I join ZDV?'"`
	Answer    string    `json:"answer" example:"There are no reasons not to join ZDV!"`
	Category  string    `gorm:"type:enum('membership', 'training', 'technology', 'misc');" json:"category" example:"membership"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy uint      `json:"created_by" example:"1293257"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy uint      `json:"updated_by" example:"1293257"`
}
