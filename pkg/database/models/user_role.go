package models

import (
	"primary-api/pkg/database"
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

func GetAllUserRolesByCID(cid uint) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("cid = ?", cid).Find(&userRoles).Error
}

func GetAllUserRolesByRoleID(roleID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("role_id = ?", roleID).Find(&userRoles).Error
}

func GetAllUserRolesByFacilityID(facilityID string) ([]UserRole, error) {
	var userRoles []UserRole
	return userRoles, database.DB.Where("facility_id = ?", facilityID).Find(&userRoles).Error
}

type Role struct {
	Role      string    `gorm:"size:4;primaryKey" json:"role" example:"ATM"`
	Name      string    `json:"name" example:"Air Traffic Manager"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func (r *Role) Create() error {
	return database.DB.Create(r).Error
}

func (r *Role) Update() error {
	return database.DB.Save(r).Error
}

func (r *Role) Delete() error {
	return database.DB.Delete(r).Error
}

func (r *Role) Get() error {
	return database.DB.Where("role = ?", r.Role).First(r).Error
}

func GetAllRoles() ([]Role, error) {
	var roles []Role
	return roles, database.DB.Find(&roles).Error
}
