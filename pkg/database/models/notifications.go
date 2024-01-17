package models

import (
	"github.com/VATUSA/primary-api/pkg/database"
	"time"
)

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

func (n *Notifications) Create() error {
	return database.DB.Create(n).Error
}

func (n *Notifications) Update() error {
	return database.DB.Save(n).Error
}

func (n *Notifications) Delete() error {
	return database.DB.Delete(n).Error
}

func (n *Notifications) Get() error {
	return database.DB.Where("id = ?", n.ID).First(n).Error
}

func GetAllNotifications() ([]Notifications, error) {
	var notifications []Notifications
	return notifications, database.DB.Find(&notifications).Error
}

func GetAllActiveNotificationsByCID(cid uint) ([]Notifications, error) {
	var notifications []Notifications
	return notifications, database.DB.Where("cid = ? AND expire_at > ?", cid, time.Now()).Find(&notifications).Error
}
