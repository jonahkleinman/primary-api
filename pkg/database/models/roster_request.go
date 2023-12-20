package models

import (
	"primary-api/pkg/database/types"
	"time"
)

type RosterRequest struct {
	ID          uint              `json:"id" gorm:"primaryKey" example:"1"`
	CID         uint              `json:"cid" example:"1293257"`
	Facility    string            `json:"requested_facility" example:"ZDV"`
	RequestType types.RequestType `gorm:"type:enum('visiting', 'transferring');"`
	Status      types.StatusType  `gorm:"type:enum('pending', 'accepted', 'rejected');"`
	Reason      string            `json:"reason" example:"I want to transfer to ZDV"`
	CreatedAt   time.Time         `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt   time.Time         `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
