package models

import "time"

type RatingChange struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID       uint      `json:"cid" example:"1293257"`
	OldRating uint      `json:"old_rating" example:"1"`
	NewRating uint      `json:"new_rating" example:"2"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy uint      `json:"created_by" example:"1293257"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
