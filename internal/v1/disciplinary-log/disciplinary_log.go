package disciplinary_log

import (
	"encoding/json"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
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
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		return err
	}
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
		return nil
	}
	return nil
}

func NewDisciplinaryLogEntryListResponse(dle []models.DisciplinaryLogEntry) []Response {
	list := []Response{}
	for _, d := range dle {
		list = append(list, *NewDisciplinaryLogEntryResponse(&d))
	}
	return list
}

func CreateDisciplinaryLogEntry(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
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
		http.Error(w, "Failed to create disciplinary log entry", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, NewDisciplinaryLogEntryResponse(dle))
}

func GetDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	render.JSON(w, r, NewDisciplinaryLogEntryResponse(dle))
}

func ListDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle, err := models.GetAllDisciplinaryLogEntries(database.DB, true)
	if err != nil {
		http.Error(w, "Failed to fetch disciplinary log entries", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, NewDisciplinaryLogEntryListResponse(dle))
}

func UpdateDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	dle.Entry = data.Entry

	if data.VATUSAOnly {
		dle.VATUSAOnly = true
	}

	if err := dle.Update(database.DB); err != nil {
		http.Error(w, "Failed to update disciplinary log entry", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, NewDisciplinaryLogEntryResponse(dle))
}

func PatchDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if data.Entry != "" {
		dle.Entry = data.Entry
	}

	if data.VATUSAOnly {
		dle.VATUSAOnly = true
	}

	if err := dle.Update(database.DB); err != nil {
		http.Error(w, "Failed to update disciplinary log entry", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, NewDisciplinaryLogEntryResponse(dle))
}

func DeleteDisciplinaryLog(w http.ResponseWriter, r *http.Request) {
	dle := GetDisciplinaryLogCtx(r)
	if err := dle.Delete(database.DB); err != nil {
		http.Error(w, "Failed to delete disciplinary log entry", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, NewDisciplinaryLogEntryResponse(dle))
}
