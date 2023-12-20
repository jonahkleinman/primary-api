package models

import (
	"time"
)

type News struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	Facility    string    `json:"facility" example:"ZDV"`
	Title       string    `json:"news" example:"DP001 Revision 3 Released"`
	Description string    `json:"answer" example:"DP001 has been revised to include new information regarding the new VATSIM Code of Conduct"`
	CreatedAt   time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy   uint      `json:"created_by" example:"1293257"`
	UpdatedAt   time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy   uint      `json:"updated_by" example:"1293257"`
}
