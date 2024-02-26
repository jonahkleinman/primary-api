package rating_change

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
	CID          uint   `json:"cid" example:"1293257" validate:"required"`
	OldRating    uint   `json:"old_rating" example:"1" validate:"required"`
	NewRating    uint   `json:"new_rating" example:"2" validate:"required"`
	CreatedByCID string `json:"created_by_cid" example:"1293257" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
	return nil
}

type Response struct {
	*models.RatingChange
}

func NewRatingChangeResponse(rc *models.RatingChange) *Response {
	return &Response{RatingChange: rc}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.RatingChange == nil {
		return errors.New("rating change not found")
	}
	return nil
}

func NewRatingChangeListResponse(rc []models.RatingChange) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range rc {
		list = append(list, NewRatingChangeResponse(&d))
	}
	return list
}

func CreateRatingChange(w http.ResponseWriter, r *http.Request) {
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

	rc := &models.RatingChange{
		CID:          data.CID,
		OldRating:    data.OldRating,
		NewRating:    data.NewRating,
		CreatedByCID: data.CreatedByCID,
	}

	if err := rc.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewRatingChangeResponse(rc))
}

func GetRatingChange(w http.ResponseWriter, r *http.Request) {
	rc := GetRatingChangeCtx(r)

	render.Render(w, r, NewRatingChangeResponse(rc))
}

func ListRatingChanges(w http.ResponseWriter, r *http.Request) {
	rc, err := models.GetAllRatingChanges(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := render.RenderList(w, r, NewRatingChangeListResponse(rc)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateRatingChange(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	rc := GetRatingChangeCtx(r)

	if !models.IsValidUser(database.DB, data.CID) {
		render.Render(w, r, utils.ErrInvalidCID)
		return
	}

	rc.CID = data.CID
	rc.OldRating = data.OldRating
	rc.NewRating = data.NewRating
	rc.CreatedByCID = data.CreatedByCID

	if err := rc.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewRatingChangeResponse(rc))
}

func PatchRatingChange(w http.ResponseWriter, r *http.Request) {
	rc := GetRatingChangeCtx(r)
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
		rc.CID = data.CID
	}
	if data.OldRating != 0 {
		rc.OldRating = data.OldRating
	}
	if data.NewRating != 0 {
		rc.NewRating = data.NewRating
	}
	if data.CreatedByCID != "" {
		rc.CreatedByCID = data.CreatedByCID
	}

	if err := rc.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewRatingChangeResponse(rc))
}

func DeleteRatingChange(w http.ResponseWriter, r *http.Request) {
	rc := GetRatingChangeCtx(r)
	if err := rc.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
