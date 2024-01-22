package models

import (
	"gorm.io/gorm"
	"time"
)

// Possible DisciplinaryLogEntry types:
// - Adding any flags
// - Changes to rating SUSPENDED

type DisciplinaryLogEntry struct {
	ID         uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID        uint      `json:"cid" example:"1293257"`
	Entry      string    `json:"entry" example:"Changed Preferred OIs to RP"`
	VATUSAOnly bool      `json:"vatusa_only" example:"true"`
	CreatedAt  time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy  string    `json:"created_by" example:"'1234567' or 'System'"`
	UpdatedAt  time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy  string    `json:"updated_by" example:"'1234567' or 'System'"`
}

func (dle *DisciplinaryLogEntry) Create(db *gorm.DB) error {
	return db.Create(dle).Error
}

func (dle *DisciplinaryLogEntry) Update(db *gorm.DB) error {
	return db.Save(dle).Error
}

func (dle *DisciplinaryLogEntry) Delete(db *gorm.DB) error {
	return db.Delete(dle).Error
}

func (dle *DisciplinaryLogEntry) Get(db *gorm.DB) error {
	return db.Where("id = ?", dle.ID).First(dle).Error
}

func GetAllDisciplinaryLogEntries(db *gorm.DB, VATUSAOnly bool) ([]DisciplinaryLogEntry, error) {
	var dle []DisciplinaryLogEntry
	return dle, db.Where("vatusa_only = ?", VATUSAOnly).Find(&dle).Error
}

func GetAllDisciplinaryLogEntriesByCID(db *gorm.DB, cid uint, VATUSAOnly bool) ([]DisciplinaryLogEntry, error) {
	var dle []DisciplinaryLogEntry
	return dle, db.Where("cid = ? AND vatusa_only = ?", cid, VATUSAOnly).Find(&dle).Error
}
