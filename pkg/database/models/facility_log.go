package models

import (
	"gorm.io/gorm"
	"time"
)

// Possible FacilityLogEntry types:
// - Changes to Facility Table
// - Changes to Faq Table (for given Facility)
// - Changes to Document Table (for given Facility)
// - Changes to Role Table (for given Facility)
// - Changes to gSuite Email (for given Facility)

type FacilityLogEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Facility  string    `json:"facility" example:"ZDV"`
	Entry     string    `json:"entry" example:"Change URL from 'denartcc.org' to 'zdvartcc.org'"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy string    `json:"created_by" example:"'1234567' or 'System'"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy string    `json:"updated_by" example:"'1234567' or 'System'"`
}

func (fle *FacilityLogEntry) Create(db *gorm.DB) error {
	return db.Create(fle).Error
}

func (fle *FacilityLogEntry) Update(db *gorm.DB) error {
	return db.Save(fle).Error
}

func (fle *FacilityLogEntry) Delete(db *gorm.DB) error {
	return db.Delete(fle).Error
}

func (fle *FacilityLogEntry) Get(db *gorm.DB) error {
	return db.Where("id = ?", fle.ID).First(fle).Error
}

func GetAllFacilityLogEntries(db *gorm.DB) ([]FacilityLogEntry, error) {
	var fle []FacilityLogEntry
	return fle, db.Find(&fle).Error
}

func GetAllFacilityLogEntriesByFacility(db *gorm.DB, facility string) ([]FacilityLogEntry, error) {
	var fle []FacilityLogEntry
	return fle, db.Where("facility = ?", facility).Find(&fle).Error
}
