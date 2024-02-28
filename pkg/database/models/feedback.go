package models

import (
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/types"
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

func (f *Feedback) Create() error {
	return database.DB.Create(f).Error
}

func (f *Feedback) Update() error {
	return database.DB.Save(f).Error
}

func (f *Feedback) Delete() error {
	return database.DB.Delete(f).Error
}

func (f *Feedback) Get() error {
	return database.DB.Where("id = ?", f.ID).First(f).Error
}

func GetAllFeedback() ([]Feedback, error) {
	var feedback []Feedback
	return feedback, database.DB.Find(&feedback).Error
}
