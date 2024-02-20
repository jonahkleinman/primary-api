package models

import (
	"gorm.io/gorm"
	"time"
)

// Expire Time can be the time of the session, or the time of the event

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	CID       uint      `json:"cid" example:"1293257"`
	Category  string    `json:"category" example:"Training"`
	Title     string    `json:"title" example:"Upcoming Training Session"`
	Body      string    `json:"body" example:"You have a training session coming up."`
	ExpireAt  time.Time `json:"expire_at" example:"2021-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (n *Notification) Create(db *gorm.DB) error {
	return db.Create(n).Error
}

func (n *Notification) Update(db *gorm.DB) error {
	return db.Save(n).Error
}

func (n *Notification) Delete(db *gorm.DB) error {
	return db.Delete(n).Error
}

func (n *Notification) Get(db *gorm.DB) error {
	return db.Where("id = ?", n.ID).First(n).Error
}

func GetAllNotifications(db *gorm.DB) ([]Notification, error) {
	var notifications []Notification
	return notifications, db.Find(&notifications).Error
}

func GetAllActiveNotificationsByCID(db *gorm.DB, cid uint) ([]Notification, error) {
	var notifications []Notification
	return notifications, db.Where("cid = ? AND expire_at > ?", cid, time.Now()).Find(&notifications).Error
}
