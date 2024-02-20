package models

import (
	"gorm.io/gorm"
	"time"
)

type RatingChange struct {
	ID           uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID          uint      `json:"cid" example:"1293257"`
	OldRating    uint      `json:"old_rating" example:"1"`
	NewRating    uint      `json:"new_rating" example:"2"`
	CreatedAt    time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedByCID string    `json:"created_by_cid" example:"1293257"`
	UpdatedAt    time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (rc *RatingChange) Create(db *gorm.DB) error {
	return db.Create(rc).Error
}

func (rc *RatingChange) Update(db *gorm.DB) error {
	return db.Save(rc).Error
}

func (rc *RatingChange) Delete(db *gorm.DB) error {
	return db.Delete(rc).Error
}

func (rc *RatingChange) Get(db *gorm.DB) error {
	return db.Where("id = ?", rc.ID).First(rc).Error
}

func GetAllRatingChanges(db *gorm.DB) ([]RatingChange, error) {
	var ratingChanges []RatingChange
	return ratingChanges, db.Find(&ratingChanges).Error
}

func GetAllRatingChangesByCID(db *gorm.DB, cid uint) ([]RatingChange, error) {
	var ratingChanges []RatingChange
	return ratingChanges, db.Where("cid = ?", cid).Find(&ratingChanges).Error
}
