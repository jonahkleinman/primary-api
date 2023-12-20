package models

import (
	"primary-api/pkg/database/types"
	"time"
)

// Feedback left for controllers will have the Facility set to the facility of the controller's position.
// Feedback left for pilots will have the Facility set to the division.
type Feedback struct {
	ID            uint                 `json:"id" gorm:"primaryKey" example:"1"`
	PilotCID      uint                 `json:"-" example:"1293257"`
	Pilot         User                 `json:"pilot" gorm:"foreignKey:PilotCID;references:CID"`
	Callsign      string               `json:"callsign" example:"DAL123"`
	ControllerCID uint                 `json:"-" example:"1293257"`
	Controller    User                 `json:"controller" gorm:"foreignKey:ControllerCID;references:CID"`
	Position      string               `json:"position" example:"DEN_I_APP"`
	Facility      string               `json:"facility" example:"ZDV"`
	Rating        types.FeedbackRating `gorm:"type:enum('unsatisfactory', 'poor', 'fair', 'good', 'excellent');" json:"rating" example:"1"`
	Notes         string               `json:"notes" example:"Raaj was the best controller I've ever flown under."`
	Status        types.StatusType     `gorm:"type:enum('pending', 'approved', 'denied');" json:"status" example:"pending"`
	Comment       string               `json:"comment" example:"Great work Raaj!"`
	CreatedAt     time.Time            `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt     time.Time            `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
