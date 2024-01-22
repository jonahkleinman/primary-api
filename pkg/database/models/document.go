package models

import (
	"github.com/VATUSA/primary-api/pkg/database/types"
	"gorm.io/gorm"
	"time"
)

type Document struct {
	ID          uint                   `json:"id" gorm:"primaryKey" example:"1"`
	Facility    string                 `json:"facility" example:"ZDV"`
	Name        string                 `json:"name" example:"DP001"`
	Description string                 `json:"description" example:"General Division Policy"`
	Category    types.DocumentCategory `gorm:"type:enum('general', 'training', 'information_technology', 'sops', 'loas', 'misc');" json:"category" example:"general"`
	URL         string                 `json:"url" example:"https://zdvartcc.org"`
	CreatedAt   time.Time              `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy   uint                   `json:"created_by" example:"1293257"`
	UpdatedAt   time.Time              `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy   uint                   `json:"updated_by" example:"1293257"`
}

func (d *Document) Create(db *gorm.DB) error {
	return db.Create(d).Error
}

func (d *Document) Update(db *gorm.DB) error {
	return db.Save(d).Error
}

func (d *Document) Delete(db *gorm.DB) error {
	return db.Delete(d).Error
}

func (d *Document) Get(db *gorm.DB) error {
	return db.Where("id = ?", d.ID).First(d).Error
}

func GetAllDocuments(db *gorm.DB) ([]Document, error) {
	var documents []Document
	return documents, db.Find(&documents).Error
}

func GetAllDocumentsByCategory(db *gorm.DB, category types.DocumentCategory) ([]Document, error) {
	var documents []Document
	return documents, db.Where("category = ?", category).Find(&documents).Error
}

func GetAllDocumentsByFacility(db *gorm.DB, facility string) ([]Document, error) {
	var documents []Document
	return documents, db.Where("facility = ?", facility).Find(&documents).Error
}

func GetAllDocumentsByFacilityAndCategory(db *gorm.DB, facility string, category types.DocumentCategory) ([]Document, error) {
	var documents []Document
	return documents, db.Where("facility = ? AND category = ?", facility, category).Find(&documents).Error
}
