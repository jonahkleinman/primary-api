package models

import (
	"github.com/VATUSA/primary-api/pkg/constants"
	"github.com/VATUSA/primary-api/pkg/database"
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID         uint             `json:"id" gorm:"primaryKey" example:"1"`
	CID        uint             `json:"cid" example:"1293257"`
	RoleID     constants.RoleID `json:"role" gorm:"type:varchar(10)" example:"ATM"`
	FacilityID string           `json:"facility_id" example:"ZDV"`
	CreatedAt  time.Time        `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt  time.Time        `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (ur *UserRole) Create() error {
	return database.DB.Create(ur).Error
}

func (ur *UserRole) Update() error {
	return database.DB.Save(ur).Error
}

func (ur *UserRole) Delete() error {
	return database.DB.Delete(ur).Error
}

func (ur *UserRole) Get() error {
	return database.DB.Where("id = ?", ur.ID).First(ur).Error
}

func GetAllUserRoles() ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Find(&userRoles).Error
}

func GetAllUserRolesByCID(db *gorm.DB, cid uint) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("cid = ?", cid).Find(&userRoles).Error
}

func GetAllUserRolesByRoleID(db *gorm.DB, roleID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("role_id = ?", roleID).Find(&userRoles).Error
}

func GetAllUserRolesByFacilityID(db *gorm.DB, facilityID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("facility_id = ?", facilityID).Find(&userRoles).Error
}

func CanModifyRole(user *User, role constants.RoleID) bool {
	if _, ok := constants.Roles[role]; !ok {
		return false
	}

	return HasRoleList(user, constants.Roles[role].RolesCanAdd)
}

func HasRoleList(user *User, roles []constants.RoleID) bool {
	for _, role := range roles {
		if HasRole(user, role) {
			return true
		}
	}
	return false
}

func HasRole(user *User, role constants.RoleID) bool {
	for _, r := range user.Roles {
		if r.RoleID == role {
			return true
		}
	}
	return false
}
