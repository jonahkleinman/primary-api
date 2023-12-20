package models

import (
	"time"
)

type UserRole struct {
	ID         uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID        uint      `json:"cid" example:"1293257"`
	RoleID     string    `json:"-" example:"ATM"`
	Role       Role      `json:"role"`
	FacilityID string    `json:"facility_id" example:"ZDV"`
	CreatedAt  time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

type Role struct {
	Role      string    `gorm:"primaryKey" json:"role" example:"ATM"`
	Name      string    `json:"name" example:"Air Traffic Manager"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
