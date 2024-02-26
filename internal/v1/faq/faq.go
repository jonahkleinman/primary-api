package faq

import (
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	Facility string `json:"facility" validate:"required,len=3"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
	Category string `json:"category" validate:"required,oneof=membership training technology misc"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.FAQ
}

func NewFAQResponse(faq *models.FAQ) *Response {
	return &Response{FAQ: faq}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.FAQ == nil {
		return nil
	}
	return nil
}

func NewFAQListResponse(faqs []models.FAQ) []render.Renderer {
	list := []render.Renderer{}
	for _, f := range faqs {
		list = append(list, NewFAQResponse(&f))
	}
	return list
}

func CreateFAQ(w http.ResponseWriter, r *http.Request) {
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

	faq := &models.FAQ{
		Facility:  data.Facility,
		Question:  data.Question,
		Answer:    data.Answer,
		Category:  data.Category,
		CreatedBy: 1,
	}

	if err := faq.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFAQResponse(faq))

}

func GetFAQ(w http.ResponseWriter, r *http.Request) {
	faq := GetFAQCtx(r)

	render.Render(w, r, NewFAQResponse(faq))
}

func ListFAQ(w http.ResponseWriter, r *http.Request) {
	faqs, err := models.GetAllFAQ(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewFAQListResponse(faqs)); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}
}

func UpdateFAQ(w http.ResponseWriter, r *http.Request) {
	faq := GetFAQCtx(r)

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

	faq.Facility = data.Facility
	faq.Question = data.Question
	faq.Answer = data.Answer
	faq.Category = data.Category

	if err := faq.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewFAQResponse(faq))
}

func PatchFAQ(w http.ResponseWriter, r *http.Request) {
	faq := GetFAQCtx(r)

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
		faq.Facility = data.Facility
	}
	if data.Question != "" {
		faq.Question = data.Question
	}
	if data.Answer != "" {
		faq.Answer = data.Answer
	}
	if data.Category != "" {
		faq.Category = data.Category
	}

	if err := faq.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewFAQResponse(faq))
}

func DeleteFAQ(w http.ResponseWriter, r *http.Request) {
	faq := GetFAQCtx(r)

	if err := faq.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
