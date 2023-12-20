package models

import (
	"primary-api/pkg/database/types"
	"time"
)

type Document struct {
	ID          uint                   `json:"id" gorm:"primaryKey" example:"1"`
	Facility    string                 `json:"facility" example:"ZDV"`
	Name        string                 `json:"name" example:"DP001"`
	Description string                 `json:"description" example:"General Division Policy"`
	Category    types.DocumentCategory `gorm:"type:enum('general', 'training', 'information_technology', 'sops', 'loas', 'misc');" json:"category" example:"general"`
	URL         string                 `json:"url" example:"https://zdvartcc.org"`
	CreatedAt   time.Time              `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy   uint                   `json:"created_by" example:"1293257"`
	UpdatedAt   time.Time              `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy   uint                   `json:"updated_by" example:"1293257"`
}
