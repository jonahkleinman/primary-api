package document

import (
	"encoding/json"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/database/types"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	Facility    string `json:"facility" example:"ZDV" validate:"required"`
	Name        string `json:"name" example:"DP001" validate:"required"`
	Description string `json:"description" example:"General Division Policy" validate:"required"`
	Category    string `json:"category" example:"general" validate:"required,oneof=general training information_technology sops loas misc"`
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
	*models.Document
}

func NewDocumentResponse(d *models.Document) *Response {
	return &Response{Document: d}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Document == nil {
		return nil
	}
	return nil
}

func NewDocumentListResponse(d []models.Document) []Response {
	list := []Response{}
	for _, doc := range d {
		list = append(list, *NewDocumentResponse(&doc))
	}
	return list
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}



	document := &models.Document{
		Facility:    data.Facility,
		Name:        data.Name,
		Description: data.Description,
		Category:    types.DocumentCategory(data.Category),
		URL:
	}
}
