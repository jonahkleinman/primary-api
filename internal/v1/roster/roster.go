package roster

import (
	"encoding/json"
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	CID        uint   `json:"cid" example:"1293257" validate:"required"`
	Facility   string `json:"facility" example:"ZDV" validate:"required"`
	OIs        string `json:"operating_initials" example:"RP" validate:"required"`
	Home       bool   `json:"home" example:"true"`
	Visiting   bool   `json:"visiting" example:"false"`
	Status     string `json:"status" example:"Active" validate:"required,oneof=active loa"` // Active, LOA
	Mentor     bool   `json:"mentor" example:"false"`
	Instructor bool   `json:"instructor" example:"false"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		return err
	}

	return nil
}

type Response struct {
	*models.Roster
}

func NewRosterResponse(r *models.Roster) *Response {
	return &Response{Roster: r}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Roster == nil {
		return errors.New("roster not found")
	}

	return nil
}

func NewRosterListResponse(r []models.Roster) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range r {
		list = append(list, NewRosterResponse(&d))
	}

	return list
}

func CreateRoster(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !data.Home && !data.Visiting {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("home and visiting cannot both be false")))
		return
	}

	if data.Home && data.Visiting {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("home and visiting cannot both be true")))
		return
	}

	roster := &models.Roster{
		CID:        data.CID,
		Facility:   data.Facility,
		OIs:        data.OIs,
		Home:       data.Home,
		Visiting:   data.Visiting,
		Status:     data.Status,
		Mentor:     data.Mentor,
		Instructor: data.Instructor,
	}

	if err := roster.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewRosterResponse(roster))
}

func GetRoster(w http.ResponseWriter, r *http.Request) {
	roster := GetRosterCtx(r)

	render.Render(w, r, NewRosterResponse(roster))
}

func ListRoster(w http.ResponseWriter, r *http.Request) {
	rosters, err := models.GetAllRosters(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.RenderList(w, r, NewRosterListResponse(rosters))
}

func UpdateRoster(w http.ResponseWriter, r *http.Request) {
	roster := GetRosterCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	roster.CID = data.CID
	roster.Facility = data.Facility
	roster.OIs = data.OIs
	roster.Home = data.Home
	roster.Visiting = data.Visiting
	roster.Status = data.Status
	roster.Mentor = data.Mentor
	roster.Instructor = data.Instructor

	if err := roster.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func DeleteRoster(w http.ResponseWriter, r *http.Request) {
	roster := GetRosterCtx(r)

	if err := roster.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
