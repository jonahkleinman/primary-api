package models

import (
	"gorm.io/gorm"
	"time"
)

type Facility struct {
	ID               string             `json:"id" gorm:"size:3;primaryKey" example:"ZDV"`
	Name             string             `json:"name" example:"Denver ARTCC"`
	URL              string             `json:"url" example:"https://zdvartcc.org"`
	FacilityLogEntry []FacilityLogEntry `json:"-" gorm:"foreignKey:Facility"`
	FAQ              []FAQ              `json:"-" gorm:"foreignKey:Facility"`
	Document         []Document         `json:"-" gorm:"foreignKey:Facility"`
	CreatedAt        time.Time          `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt        time.Time          `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (f *Facility) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *Facility) Update(db *gorm.DB) error {
	return db.Save(f).Error
}

func (f *Facility) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *Facility) Get(db *gorm.DB) error {
	return db.Where("id = ?", f.ID).First(f).Error
}

func IsValidFacility(db *gorm.DB, id string) bool {
	var f Facility
	return db.Where("id = ?", id).First(&f).Error == nil
}

func GetAllFacilities(db *gorm.DB) ([]Facility, error) {
	var facilities []Facility
	return facilities, db.Find(&facilities).Error
}
