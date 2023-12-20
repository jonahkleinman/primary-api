package models

import "time"

type Facility struct {
	ID        string    `json:"id" gorm:"primaryKey" example:"ZDV"`
	Name      string    `json:"name" example:"Denver ARTCC"`
	URL       string    `json:"url" example:"https://zdvartcc.org"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
