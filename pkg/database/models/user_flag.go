package models

import (
	"gorm.io/gorm"
	"time"
)

type Flag struct {
	ID                       uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID                      uint      `json:"cid" example:"1293257"`
	NoStaffRole              bool      `json:"no_staff_role" example:"false"`
	NoStaffLogEntryID        uint      `json:"no_staff_log_entry_id" example:"1"`
	NoVisiting               bool      `json:"no_visiting" example:"false"`
	NoVisitingLogEntryID     uint      `json:"no_visiting_log_entry_id" example:"1"`
	NoTransferring           bool      `json:"no_transferring" example:"false"`
	NoTransferringLogEntryID uint      `json:"no_transferring_log_entry_id" example:"1"`
	NoTraining               bool      `json:"no_training" example:"false"`
	NoTrainingLogEntryID     uint      `json:"no_training_log_entry_id" example:"1"`
	CreatedAt                time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt                time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (f *Flag) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *Flag) Update(db *gorm.DB) error {
	return db.Save(f).Error
}

func (f *Flag) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *Flag) Get(db *gorm.DB) error {
	return db.Where("id = ?", f.ID).First(f).Error
}

func GetAllFlags(db *gorm.DB) ([]Flag, error) {
	var flags []Flag
	return flags, db.Find(&flags).Error
}

func GetAllFlagsByCID(db *gorm.DB, cid uint) ([]Flag, error) {
	var flags []Flag
	return flags, db.Where("cid = ?", cid).Find(&flags).Error
}
