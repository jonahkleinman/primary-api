package models

import (
	"github.com/VATUSA/primary-api/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID         uint             `json:"id" gorm:"primaryKey" example:"1"`
	CID        uint             `json:"cid" example:"1293257"`
	RoleID     constants.RoleID `json:"role" example:"ATM"`
	FacilityID string           `json:"facility_id" example:"ZDV"`
	CreatedAt  time.Time        `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt  time.Time        `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (ur *UserRole) Create(db *gorm.DB) error {
	return db.Create(ur).Error
}

func (ur *UserRole) Update(db *gorm.DB) error {
	return db.Save(ur).Error
}

func (ur *UserRole) Delete(db *gorm.DB) error {
	return db.Delete(ur).Error
}

func (ur *UserRole) Get(db *gorm.DB) error {
	return db.Where("id = ?", ur.ID).First(ur).Error
}

func GetAllUserRoles(db *gorm.DB) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, db.Find(&userRoles).Error
}

func GetAllUserRolesByCID(db *gorm.DB, cid uint) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, db.Where("cid = ?", cid).Find(&userRoles).Error
}

func GetAllUserRolesByRoleID(db *gorm.DB, roleID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, db.Where("role_id = ?", roleID).Find(&userRoles).Error
}

func GetAllUserRolesByFacilityID(db *gorm.DB, facilityID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, db.Where("facility_id = ?", facilityID).Find(&userRoles).Error
}
