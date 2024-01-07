package models

import (
	"primary-api/pkg/database"
	"time"
)

type News struct {
	ID           uint      `json:"id" gorm:"primaryKey" example:"1"`
	Facility     string    `json:"facility" example:"ZDV"`
	Title        string    `json:"news" example:"DP001 Revision 3 Released"`
	Description  string    `json:"answer" example:"DP001 has been revised to include new information regarding the new VATSIM Code of Conduct"`
	CreatedAt    time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedByCID uint      `json:"created_by_cid" example:"1293257"`
	UpdatedAt    time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedByCID uint      `json:"updated_by_cid" example:"1293257"`
}

func (n *News) Create() error {
	return database.DB.Create(n).Error
}

func (n *News) Update() error {
	return database.DB.Save(n).Error
}

func (n *News) Delete() error {
	return database.DB.Delete(n).Error
}

func (n *News) Get() error {
	return database.DB.Where("id = ?", n.ID).First(n).Error
}

func GetAllNews() ([]News, error) {
	var news []News
	return news, database.DB.Find(&news).Error
}
