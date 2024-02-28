package news

import (
	"errors"
	"fmt"
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

// CreateNews godoc
// @Summary Create a new news entry
// @Description Create a new news entry
// @Tags news
// @Accept  json
// @Produce  json
// @Param news body Request true "News Entry"
// @Success 201 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /news [post]
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

	if !models.IsValidFacility(data.Facility) {
		render.Render(w, r, utils.ErrInvalidFacility)
		return
	}

	news := &models.News{
		Facility:    data.Facility,
		Title:       data.Title,
		Description: data.Description,
		CreatedBy:   "System",
	}

	if err := news.Create(); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewNewsResponse(news))
}

// GetNews godoc
// @Summary Get a news entry
// @Description Get a news entry
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path string true "News ID"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /news/{id} [get]
func GetNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)
	render.Render(w, r, NewNewsResponse(news))
}

// ListNews godoc
// @Summary List all news entries
// @Description List all news entries
// @Tags news
// @Accept  json
// @Produce  json
// @Success 200 {object} []Response
// @Failure 422 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /news [get]
func ListNews(w http.ResponseWriter, r *http.Request) {
	news, err := models.GetAllNews()
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewNewsListResponse(news)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// UpdateNews godoc
// @Summary Update a news entry
// @Description Update a news entry
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path string true "News ID"
// @Param news body Request true "News Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /news/{id} [put]
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

	if !models.IsValidFacility(req.Facility) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid facility")))
		return
	}

	news.Facility = req.Facility
	news.Title = req.Title
	news.Description = req.Description
	news.UpdatedBy = "System"

	if err := news.Update(); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

// PatchNews godoc
// @Summary Patch a news entry
// @Description Patch a news entry
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path string true "News ID"
// @Param news body Request true "News Entry"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrResponse
// @Failure 404 {object} utils.ErrResponse
// @Failure 500 {object} utils.ErrResponse
// @Router /news/{id} [patch]
func PatchNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if req.Facility != "" {
		if !models.IsValidFacility(req.Facility) {
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

	news.UpdatedBy = "System"

	if err := news.Update(); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNewsResponse(news))
}

// DeleteNews godoc
// @Summary Delete a news entry
// @Description Delete a news entry
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path string true "News ID"
// @Success 204
// @Failure 500 {object} utils.ErrResponse
// @Router /news/{id} [delete]
func DeleteNews(w http.ResponseWriter, r *http.Request) {
	news := GetNewsCtx(r)
	if err := news.Delete(); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
