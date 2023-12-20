package models

import "time"

// Expire Time can be the time of the session, or the time of the event

type Notifications struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID       uint      `json:"cid" example:"1293257"`
	Category  string    `json:"category" example:"Training"`
	Title     string    `json:"title" example:"Upcoming Training Session"`
	Body      string    `json:"body" example:"You have a training session coming up."`
	ExpireAt  time.Time `json:"expire_at" example:"2021-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
