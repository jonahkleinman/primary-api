package document

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/database/types"
	"github.com/VATUSA/primary-api/pkg/storage"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
	"path"
	"strings"
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
		return errors.New("missing required document")
	}
	return nil
}

func NewDocumentListResponse(d []models.Document) []render.Renderer {
	list := []render.Renderer{}
	for _, doc := range d {
		list = append(list, NewDocumentResponse(&doc))
	}
	return list
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	currentDocs, err := models.GetAllDocumentsByFacilityAndCategory(database.DB, data.Facility, types.DocumentCategory(data.Category))
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	for _, doc := range currentDocs {
		if doc.Name == data.Name {
			render.Render(w, r, utils.ErrInvalidRequest(fmt.Errorf("document with name %s already exists", data.Name)))
			return
		}
	}

	// Read the file from the request
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}

	extension := path.Ext(fileHeader.Filename)
	directory := path.Join(data.Facility, data.Category)
	filename := strings.ReplaceAll(data.Name, " ", "-") + extension

	// Put the file in the S3 bucket
	if err := storage.PublicBucket.Upload(directory, filename, file); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	document := &models.Document{
		Facility:    data.Facility,
		Name:        data.Name,
		Description: data.Description,
		Category:    types.DocumentCategory(data.Category),
		URL:         path.Join("https://cdn.vatusa.local", directory, filename),
	}

	if err := document.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)

		err := storage.PublicBucket.Delete(directory, filename)
		if err != nil {
			fmt.Println("Error deleting file from S3:", err)
		}
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewDocumentResponse(document))
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	doc := GetDocumentCtx(r)
	render.Render(w, r, NewDocumentResponse(doc))
}

func ListDocuments(w http.ResponseWriter, r *http.Request) {
	docs, err := models.GetAllDocuments(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewDocumentListResponse(docs)); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}
}

func ListDocumentsByFac(w http.ResponseWriter, r *http.Request) {
	facId := chi.URLParam(r, "Facility")
	if facId == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	docs, err := models.GetAllDocumentsByFacility(database.DB, facId)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewDocumentListResponse(docs)); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return

	}
}

func ListDocumentsByFacByCat(w http.ResponseWriter, r *http.Request) {
	facId := chi.URLParam(r, "Facility")
	cat := chi.URLParam(r, "Category")
	if facId == "" || cat == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	docs, err := models.GetAllDocumentsByFacilityAndCategory(database.DB, facId, types.DocumentCategory(cat))
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewDocumentListResponse(docs)); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	doc := GetDocumentCtx(r)

	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	// TODO - update the file in the S3 bucket

	doc.Facility = data.Facility
	doc.Name = data.Name
	doc.Description = data.Description
	doc.Category = types.DocumentCategory(data.Category)

	if err := doc.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewDocumentResponse(doc))
}

func PatchDocument(w http.ResponseWriter, r *http.Request) {
	doc := GetDocumentCtx(r)

	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	// TODO - update the file in the S3 bucket

	if data.Facility != "" {
		doc.Facility = data.Facility
	}
	if data.Name != "" {
		doc.Name = data.Name
	}
	if data.Description != "" {
		doc.Description = data.Description
	}
	if data.Category != "" {
		doc.Category = types.DocumentCategory(data.Category)
	}

	if err := doc.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewDocumentResponse(doc))
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	doc := GetDocumentCtx(r)

	// Delete the file from the S3 bucket
	directory := path.Join(doc.Facility, string(doc.Category))
	filename := strings.ReplaceAll(doc.Name, " ", "-") + path.Ext(doc.URL)
	if err := storage.PublicBucket.Delete(directory, filename); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := doc.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)

}
