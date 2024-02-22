package facility_log

import (
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	Facility string `json:"facility" example:"ZDV" validate:"required,len=3"`
	Entry    string `json:"entry" example:"Changed Preferred OIs to RP" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.FacilityLogEntry
}

func NewFacilityLogEntryResponse(fle *models.FacilityLogEntry) *Response {
	return &Response{FacilityLogEntry: fle}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.FacilityLogEntry == nil {
		return errors.New("facility log entry not found")
	}
	return nil
}

func NewFacilityLogEntryListResponse(fle []models.FacilityLogEntry) []render.Renderer {
	list := []render.Renderer{}
	for _, f := range fle {
		list = append(list, NewFacilityLogEntryResponse(&f))
	}
	return list
}

func CreateFacilityLogEntry(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidFacility(database.DB, data.Facility) {
		render.Render(w, r, utils.ErrInvalidFacility)
		return
	}

	fle := &models.FacilityLogEntry{
		Facility:  data.Facility,
		Entry:     data.Entry,
		CreatedBy: "System",
	}

	if err := fle.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFacilityLogEntryResponse(fle))
}

func GetFacilityLog(w http.ResponseWriter, r *http.Request) {
	fle := GetFacilityLogCtx(r)

	render.Render(w, r, NewFacilityLogEntryResponse(fle))
}

func ListFacilityLog(w http.ResponseWriter, r *http.Request) {
	fle, err := models.GetAllFacilityLogEntries(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewFacilityLogEntryListResponse(fle)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateFacilityLog(w http.ResponseWriter, r *http.Request) {
	fle := GetFacilityLogCtx(r)

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidFacility(database.DB, data.Facility) {
		render.Render(w, r, utils.ErrInvalidFacility)
		return
	}

	fle.Facility = data.Facility
	fle.Entry = data.Entry

	if err := fle.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewFacilityLogEntryResponse(fle))
}

func PatchFacilityLog(w http.ResponseWriter, r *http.Request) {
	fle := GetFacilityLogCtx(r)

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if data.Facility != "" {
		if !models.IsValidFacility(database.DB, data.Facility) {
			render.Render(w, r, utils.ErrInvalidFacility)
			return
		}
		fle.Facility = data.Facility
	}
	if data.Entry != "" {
		fle.Entry = data.Entry

	}

	if err := fle.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewFacilityLogEntryResponse(fle))
}

func DeleteFacilityLog(w http.ResponseWriter, r *http.Request) {
	fle := GetFacilityLogCtx(r)

	if err := fle.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
