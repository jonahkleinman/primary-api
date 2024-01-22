package models

import (
	"github.com/VATUSA/primary-api/pkg/database/types"
	"gorm.io/gorm"
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

func (rr *RosterRequest) Create(db *gorm.DB) error {
	return db.Create(rr).Error
}

func (rr *RosterRequest) Update(db *gorm.DB) error {
	return db.Save(rr).Error
}

func (rr *RosterRequest) Delete(db *gorm.DB) error {
	return db.Delete(rr).Error
}

func (rr *RosterRequest) Get(db *gorm.DB) error {
	return db.Where("id = ?", rr.ID).First(rr).Error
}

func GetAllRosterRequests(db *gorm.DB) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Find(&rosterRequests).Error
}

func GetAllRosterRequestsByCID(db *gorm.DB, cid uint) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("cid = ?", cid).Find(&rosterRequests).Error
}

func GetAllRosterRequestsByFacility(db *gorm.DB, facility string) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("requested_facility = ?", facility).Find(&rosterRequests).Error
}

func GetAllPendingVisitingRequestsByCID(db *gorm.DB, cid uint) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("cid = ? AND request_type = ? AND status = ?", cid, types.Visiting, types.Pending).Find(&rosterRequests).Error
}

func GetAllPendingTransferringRequestsByCID(db *gorm.DB, cid uint) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("cid = ? AND request_type = ? AND status = ?", cid, types.Transferring, types.Pending).Find(&rosterRequests).Error
}

func GetAllPendingVisitingRequestsByFacility(db *gorm.DB, facility string) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("requested_facility = ? AND request_type = ? AND status = ?", facility, types.Visiting, types.Pending).Find(&rosterRequests).Error
}

func GetAllPendingTransferringRequestsByFacility(db *gorm.DB, facility string) ([]RosterRequest, error) {
	var rosterRequests []RosterRequest
	return rosterRequests, db.Where("requested_facility = ? AND request_type = ? AND status = ?", facility, types.Transferring, types.Pending).Find(&rosterRequests).Error
}
