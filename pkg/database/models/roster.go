package models

import (
	"github.com/VATUSA/primary-api/pkg/database"
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

func (r *Roster) Create() error {
	return database.DB.Create(r).Error
}

func (r *Roster) Update() error {
	return database.DB.Save(r).Error
}

func (r *Roster) Delete() error {
	return database.DB.Delete(r).Error
}

func (r *Roster) Get() error {
	return database.DB.Where("id = ?", r.ID).First(r).Error
}

func GetAllRosters() ([]Roster, error) {
	var rosters []Roster
	return rosters, database.DB.Find(&rosters).Error
}

func GetAllRostersByCID(cid uint) ([]Roster, error) {
	var rosters []Roster
	return rosters, database.DB.Where("cid = ?", cid).Find(&rosters).Error
}

func GetAllRostersByFacility(facility string) ([]Roster, error) {
	var rosters []Roster
	return rosters, database.DB.Where("facility = ?", facility).Find(&rosters).Error
}
