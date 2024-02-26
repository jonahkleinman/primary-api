package roster_request

import (
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/database/types"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	CID         uint              `json:"cid" example:"1293257" validate:"required"`
	Facility    string            `json:"requested_facility" example:"ZDV" validate:"required,len=3"`
	RequestType types.RequestType `json:"request_type" example:"visiting" validate:"required,oneof=visiting transferring"`
	Status      types.StatusType  `json:"status" example:"pending" validate:"required,oneof=pending accepted rejected"`
	Reason      string            `json:"reason" example:"I want to transfer to ZDV" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.RosterRequest
}

func NewRosterRequestResponse(r *models.RosterRequest) *Response {
	return &Response{RosterRequest: r}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.RosterRequest == nil {
		return errors.New("roster request not found")
	}
	return nil
}

func NewRosterRequestListResponse(r []models.RosterRequest) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range r {
		list = append(list, NewRosterRequestResponse(&d))
	}
	return list
}

func CreateRosterRequest(w http.ResponseWriter, r *http.Request) {
	req := &Request{}
	if err := req.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidUser(database.DB, req.CID) {
		render.Render(w, r, utils.ErrInvalidCID)
		return
	}

	if !models.IsValidFacility(database.DB, req.Facility) {
		render.Render(w, r, utils.ErrInvalidFacility)
		return
	}

	rosterRequest := &models.RosterRequest{
		CID:         req.CID,
		Facility:    req.Facility,
		RequestType: req.RequestType,
		Status:      req.Status,
		Reason:      req.Reason,
	}

	if err := rosterRequest.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewRosterRequestResponse(rosterRequest))
}

func GetRosterRequest(w http.ResponseWriter, r *http.Request) {
	rosterRequest := GetRosterRequestCtx(r)

	render.Render(w, r, NewRosterRequestResponse(rosterRequest))
}

func ListRosterRequest(w http.ResponseWriter, r *http.Request) {
	rosterRequests, err := models.GetAllRosterRequests(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := render.RenderList(w, r, NewRosterRequestListResponse(rosterRequests)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateRosterRequest(w http.ResponseWriter, r *http.Request) {
	req := GetRosterRequestCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidUser(database.DB, req.CID) {
		render.Render(w, r, utils.ErrInvalidCID)
		return
	}

	if !models.IsValidFacility(database.DB, req.Facility) {
		render.Render(w, r, utils.ErrInvalidFacility)
		return
	}

	if req.Status == types.Pending && data.Status == types.Accepted {
		roster := &models.Roster{
			CID:        data.CID,
			Facility:   data.Facility,
			OIs:        "",
			Home:       false,
			Visiting:   false,
			Status:     "Active",
			Mentor:     false,
			Instructor: false,
		}

		if data.RequestType == types.Visiting {
			roster.Visiting = true
		} else {
			roster.Home = true
		}

		if err := roster.Create(database.DB); err != nil {
			render.Render(w, r, utils.ErrInvalidRequest(err))
			return
		}
	}

	req.CID = data.CID
	req.Facility = data.Facility
	req.RequestType = data.RequestType
	req.Status = data.Status
	req.Reason = data.Reason

	if err := req.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewRosterRequestResponse(req))
}

func DeleteRosterRequest(w http.ResponseWriter, r *http.Request) {
	req := GetRosterRequestCtx(r)
	if err := req.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
