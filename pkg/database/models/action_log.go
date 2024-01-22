package models

import (
	"gorm.io/gorm"
	"time"
)

// Possible ActionLogEntries types:
// - Changes to User Table
// - Changes to Roster Table (for given User)
// - Changes to User Role Table (for given User)
// - Changes to Rating Changes Table (for given User)

type ActionLogEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID       uint      `json:"cid" example:"1293257"`
	Entry     string    `json:"entry" example:"Changed Preferred OIs to RP"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy string    `json:"created_by" example:"'1234567' or 'System'"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy string    `json:"updated_by" example:"'1234567' or 'System'"`
}

func (ale *ActionLogEntry) Create(db *gorm.DB) error {
	return db.Create(ale).Error
}

func (ale *ActionLogEntry) Update(db *gorm.DB) error {
	return db.Save(ale).Error
}

func (ale *ActionLogEntry) Delete(db *gorm.DB) error {
	return db.Delete(ale).Error
}

func (ale *ActionLogEntry) Get(db *gorm.DB) error {
	return db.Where("id = ?", ale.ID).First(ale).Error
}

func GetAllActionLogEntries(db *gorm.DB) ([]ActionLogEntry, error) {
	var ale []ActionLogEntry
	return ale, db.Find(&ale).Error
}

func GetAllActionLogEntriesByCID(db *gorm.DB, cid uint) ([]ActionLogEntry, error) {
	var ale []ActionLogEntry
	return ale, db.Where("cid = ?", cid).Find(&ale).Error
}
