package models

import "time"

type User struct {
	CID              uint      `gorm:"primaryKey" json:"cid" example:"1293257"`
	FirstName        string    `json:"first_name" example:"Raaj"`
	LastName         string    `json:"last_name" example:"Patel"`
	PreferredName    string    `json:"preferred_name" example:"Raaj"`
	PrefNameEnabled  bool      `json:"pref_name_enabled" example:"true"`
	Email            string    `json:"email" example:"vatusa6@vatusa.net"`
	PreferredOIs     string    `json:"preferred_ois" example:"RP"`
	PilotRating      uint      `json:"pilot_rating" example:"1"`
	ControllerRating uint      `json:"controller_rating" example:"1"`
	DiscordID        string    `json:"discord_id" example:"1234567890"`
	LastLogin        time.Time `json:"last_login" example:"2021-01-01T00:00:00Z"`
	LastCertSync     time.Time `json:"last_cert_sync" example:"2021-01-01T00:00:00Z"`
	CreatedAt        time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt        time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
