package action_log

import (
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	CID   uint   `json:"cid" example:"1293257" validate:"required"`
	Entry string `json:"entry" example:"Changed Preferred OIs to RP" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.ActionLogEntry
}

func NewActionLogEntryResponse(ale *models.ActionLogEntry) *Response {
	return &Response{ActionLogEntry: ale}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.ActionLogEntry == nil {
		return nil
	}
	return nil
}

func NewActionLogEntryListResponse(ale []models.ActionLogEntry) []render.Renderer {
	list := []render.Renderer{}
	for _, a := range ale {
		list = append(list, NewActionLogEntryResponse(&a))
	}
	return list
}

// CreateActionLogEntry godoc
// @Summary Create a new action log entry
// @Description Create a new action log entry
// @Tags action-log
// @Accept  json
// @Produce  json
// @Param action_log body Request true "Action Log Entry"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log [post]
func CreateActionLogEntry(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
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

	ale := &models.ActionLogEntry{
		CID:       data.CID,
		Entry:     data.Entry,
		CreatedBy: "System",
	}

	if err := ale.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActionLogEntryResponse(ale))
}

// GetActionLog godoc
// @Summary Get an action log entry
// @Description Get an action log entry
// @Tags action-log
// @Accept  json
// @Produce  json
// @Param id path int true "Action Log Entry ID"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log/{id} [get]
func GetActionLog(w http.ResponseWriter, r *http.Request) {
	ale := GetActionLogCtx(r)

	render.Render(w, r, NewActionLogEntryResponse(ale))
}

// ListActionLog godoc
// @Summary List all action log entries
// @Description List all action log entries
// @Tags action-log
// @Accept  json
// @Produce  json
// @Success 200 {object} []Response
// @Failure 422 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log [get]
func ListActionLog(w http.ResponseWriter, r *http.Request) {
	ale, err := models.GetAllActionLogEntries(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewActionLogEntryListResponse(ale)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// UpdateActionLog godoc
// @Summary Update an action log entry
// @Description Update an action log entry
// @Tags action-log
// @Accept  json
// @Produce  json
// @Param id path int true "Action Log Entry ID"
// @Param action_log body Request true "Action Log Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log/{id} [put]
func UpdateActionLog(w http.ResponseWriter, r *http.Request) {
	ale := GetActionLogCtx(r)

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
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

	ale.CID = data.CID
	ale.Entry = data.Entry
	ale.UpdatedBy = "System"

	if err := ale.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewActionLogEntryResponse(ale))
}

// PatchActionLog godoc
// @Summary Patch an action log entry
// @Description Patch an action log entry
// @Tags action-log
// @Accept  json
// @Produce  json
// @Param id path int true "Action Log Entry ID"
// @Param action_log body Request true "Action Log Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log/{id} [patch]
func PatchActionLog(w http.ResponseWriter, r *http.Request) {
	ale := GetActionLogCtx(r)

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if data.CID != 0 {
		if !models.IsValidUser(database.DB, data.CID) {
			render.Render(w, r, utils.ErrInvalidCID)
			return
		}
		ale.CID = data.CID
	}
	if data.Entry != "" {
		ale.Entry = data.Entry
	}
	ale.UpdatedBy = "System"

	if err := ale.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewActionLogEntryResponse(ale))
}

// DeleteActionLog godoc
// @Summary Delete an action log entry
// @Description Delete an action log entry
// @Tags action-log
// @Accept  json
// @Produce  json
// @Param id path int true "Action Log Entry ID"
// @Success 204
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /action-log/{id} [delete]
func DeleteActionLog(w http.ResponseWriter, r *http.Request) {
	ale := GetActionLogCtx(r)

	if err := ale.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
