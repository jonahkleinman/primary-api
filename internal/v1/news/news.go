package news

import (
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

// TODO - Add UpdatedByCID functionality

type Request struct {
	*models.News
}

func (req *Request) Bind(r *http.Request) error {
	if req.News == nil {
		return errors.New("missing required News fields")
	}

	missingFields := []string{}
	if req.Facility == "" {
		missingFields = append(missingFields, "facility")
	}
	if req.Title == "" {
		missingFields = append(missingFields, "title")
	}
	if req.Description == "" {
		missingFields = append(missingFields, "description")
	}

	if len(missingFields) > 0 {
		return errors.New("missing required fields: " + strings.Join(missingFields, ", "))
	}

	return nil
}

func (req *Request) BindPartial(r *http.Request) error {
	if req.News == nil {
		return errors.New("missing required News fields")
	}

	return nil
}

type Response struct {
	*models.News
}

func NewNewsResponse(news *models.News) *Response {
	resp := &Response{News: news}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
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
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	news := data.News
	if err := news.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewNewsResponse(news))
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	news := r.Context().Value("news").(*models.News)

	render.Render(w, r, NewNewsResponse(news))
}

func ListNews(w http.ResponseWriter, r *http.Request) {
	news, err := models.GetAllNews(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewNewsListResponse(news)); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	news := r.Context().Value("news").(*models.News)

	data := &Request{News: news}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}
	news = data.News
	if err := news.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	news := r.Context().Value("news").(*models.News)

	data := &Request{News: news}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if data.Facility != "" {
		news.Facility = data.Facility
	}
	if data.Title != "" {
		news.Title = data.Title
	}
	if data.Description != "" {
		news.Description = data.Description
	}

	if err := news.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

func PatchNews(w http.ResponseWriter, r *http.Request) {
	news := r.Context().Value("news").(*models.News)

	data := &Request{News: news}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}
	news = data.News
	if err := news.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	news := r.Context().Value("news").(*models.News)

	if err := news.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
