package news

import (
	"errors"
	"fmt"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	Facility    string `json:"facility" example:"ZDV" validate:"required,len=3"`
	Title       string `json:"title" example:"DP001 Revision 3 Released" validate:"required"`
	Description string `json:"description" example:"DP001 has been revised to include new information regarding the new VATSIM Code of Conduct" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.News
}

func NewNewsResponse(news *models.News) *Response {
	return &Response{News: news}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.News == nil {
		return errors.New("missing required news")
	}
	return nil
}

func NewNewsListResponse(news []models.News) []render.Renderer {
	list := []render.Renderer{}
	for _, n := range news {
		list = append(list, NewNewsResponse(&n))
	}
	return list
}

func CreateNews(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		fmt.Println(r.Body)
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidFacility(database.DB, data.Facility) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid facility")))
		return
	}

	news := &models.News{
		Facility:    data.Facility,
		Title:       data.Title,
		Description: data.Description,
		CreatedBy:   "System",
	}

	if err := news.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewNewsResponse(news))
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)
	render.Render(w, r, NewNewsResponse(news))
}

func ListNews(w http.ResponseWriter, r *http.Request) {
	news, err := models.GetAllNews(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewNewsListResponse(news)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidFacility(database.DB, req.Facility) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid facility")))
		return
	}

	news.Facility = req.Facility
	news.Title = req.Title
	news.Description = req.Description

	// 1 is the vatusa user
	news.UpdatedByCID = 1

	if err := news.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

func PatchNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if req.Facility != "" {
		if !models.IsValidFacility(database.DB, req.Facility) {
			render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid facility")))
			return
		}

		news.Facility = req.Facility
	}
	if req.Title != "" {
		news.Title = req.Title
	}
	if req.Description != "" {
		news.Description = req.Description
	}

	// 1 is the vatusa user
	news.UpdatedByCID = 1

	if err := news.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)
	if err := news.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
