package models

import (
	"gorm.io/gorm"
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

func (f *FAQ) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *FAQ) Update(db *gorm.DB) error {
	return db.Save(f).Error
}

func (f *FAQ) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *FAQ) Get(db *gorm.DB) error {
	return db.Where("id = ?", f.ID).First(f).Error
}

func GetAllFAQ(db *gorm.DB) ([]FAQ, error) {
	var faq []FAQ
	return faq, db.Find(&faq).Error
}

func GetAllFAQByCategory(db *gorm.DB, category string) ([]FAQ, error) {
	var faq []FAQ
	return faq, db.Where("category = ?", category).Find(&faq).Error
}

func GetAllFAQByFacility(db *gorm.DB, facility string) ([]FAQ, error) {
	var faq []FAQ
	return faq, db.Where("facility = ?", facility).Find(&faq).Error
}
