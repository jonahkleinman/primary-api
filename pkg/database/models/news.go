package models

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	Facility    string    `json:"facility" example:"ZDV"`
	Title       string    `json:"news" example:"DP001 Revision 3 Released"`
	Description string    `json:"answer" example:"DP001 has been revised to include new information regarding the new VATSIM Code of Conduct"`
	CreatedAt   time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	CreatedBy   string    `json:"created_by" example:"'1293257' or 'System'"`
	UpdatedAt   time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	UpdatedBy   string    `json:"updated_by" example:"1293257"`
}

func (n *News) Create(db *gorm.DB) error {
	return db.Create(n).Error
}

func (n *News) Update(db *gorm.DB) error {
	return db.Save(n).Error
}

func (n *News) Delete(db *gorm.DB) error {
	return db.Delete(n).Error
}

func (n *News) Get(db *gorm.DB) error {
	return db.Where("id = ?", n.ID).First(n).Error
}

func GetAllNews(db *gorm.DB) ([]News, error) {
	var news []News
	return news, db.Find(&news).Error
}
