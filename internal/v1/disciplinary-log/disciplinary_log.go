package disciplinary_log

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
	CID        uint   `json:"cid" example:"1293257" validate:"required"`
	Entry      string `json:"entry" example:"Changed Preferred OIs to RP" validate:"required"`
	VATUSAOnly bool   `json:"vatusa_only" example:"true"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.DisciplinaryLogEntry
}

func NewDisciplinaryLogEntryResponse(dle *models.DisciplinaryLogEntry) *Response {
	return &Response{DisciplinaryLogEntry: dle}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.DisciplinaryLogEntry == nil {
		return errors.New("disciplinary log entry not found")
	}
	return nil
}

func NewDisciplinaryLogEntryListResponse(dle []models.DisciplinaryLogEntry) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range dle {
		list = append(list, NewDisciplinaryLogEntryResponse(&d))
	}
	return list
}

// CreateDisciplinaryLogEntry godoc
// @Summary Create a new disciplinary log entry
// @Description Create a new disciplinary log entry
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Param disciplinary_log body Request true "Disciplinary Log Entry"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log [post]
func CreateDisciplinaryLogEntry(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidUser(database.DB, data.CID) {
		render.Render(w, r, utils.ErrInvalidCID)
		return
	}

	dle := &models.DisciplinaryLogEntry{
		CID:       data.CID,
		Entry:     data.Entry,
		CreatedBy: "System",
	}

	if data.VATUSAOnly {
		dle.VATUSAOnly = true
	} else {
		dle.VATUSAOnly = false
	}

	if err := dle.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewDisciplinaryLogEntryResponse(dle))
}

// GetDisciplinaryLog godoc
// @Summary Get a disciplinary log entry
// @Description Get a disciplinary log entry
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Param id path string true "Disciplinary Log Entry ID"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log/{id} [get]
func GetDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	render.Render(w, r, NewDisciplinaryLogEntryResponse(dle))
}

// ListDisciplinaryLog godoc
// @Summary List all disciplinary log entries
// @Description List all disciplinary log entries
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Success 200 {object} []Response
// @Failure 422 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log [get]
func ListDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle, err := models.GetAllDisciplinaryLogEntries(database.DB, true)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewDisciplinaryLogEntryListResponse(dle)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// UpdateDisciplinaryLog godoc
// @Summary Update a disciplinary log entry
// @Description Update a disciplinary log entry
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Param id path string true "Disciplinary Log Entry ID"
// @Param disciplinary_log body Request true "Disciplinary Log Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log/{id} [put]
func UpdateDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidUser(database.DB, data.CID) {
		render.Render(w, r, utils.ErrInvalidCID)
		return
	}

	dle.CID = data.CID
	dle.Entry = data.Entry

	if data.VATUSAOnly {
		dle.VATUSAOnly = true
	}

	if err := dle.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewDisciplinaryLogEntryResponse(dle))
}

// PatchDisciplinaryLog godoc
// @Summary Patch a disciplinary log entry
// @Description Patch a disciplinary log entry
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Param id path string true "Disciplinary Log Entry ID"
// @Param disciplinary_log body Request true "Disciplinary Log Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log/{id} [patch]
func PatchDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if data.CID != 0 {
		if !models.IsValidUser(database.DB, data.CID) {
			render.Render(w, r, utils.ErrInvalidCID)
			return
		}
		dle.CID = data.CID
	}

	if data.Entry != "" {
		dle.Entry = data.Entry
	}

	if data.VATUSAOnly {
		dle.VATUSAOnly = true
	}

	if err := dle.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewDisciplinaryLogEntryResponse(dle))
}

// DeleteDisciplinaryLog godoc
// @Summary Delete a disciplinary log entry
// @Description Delete a disciplinary log entry
// @Tags disciplinary-log
// @Accept  json
// @Produce  json
// @Param id path string true "Disciplinary Log Entry ID"
// @Success 204
// @Failure 500 {object} utils.ErrResponse
// @Router /disciplinary-log/{id} [delete]
func DeleteDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	if err := dle.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
