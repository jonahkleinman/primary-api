package notification

import (
	"encoding/json"
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Request struct {
	CID      uint   `json:"cid" example:"1293257" validate:"required"`
	Category string `json:"category" example:"Training" validate:"required"`
	Title    string `json:"title" example:"Upcoming Training Session" validate:"required"`
	Body     string `json:"body" example:"You have a training session coming up." validate:"required"`
	ExpireAt string `json:"expire_at" example:"2021-01-01T00:00:00Z" validate:"required"`
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
	*models.Notification
}

func NewNotificationResponse(n *models.Notification) *Response {
	return &Response{Notification: n}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Notification == nil {
		return errors.New("notification not found")
	}
	return nil
}

func NewNotificationListResponse(n []models.Notification) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range n {
		list = append(list, NewNotificationResponse(&d))
	}
	return list
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	expireAt, err := http.ParseTime(data.ExpireAt)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	// Make sure expireAt is in the future
	if expireAt.Before(time.Now()) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("expire_at must be in the future")))
		return
	}

	n := &models.Notification{
		CID:      data.CID,
		Category: data.Category,
		Title:    data.Title,
		Body:     data.Body,
		ExpireAt: expireAt,
	}

	if err := n.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewNotificationResponse(n))
}

func GetNotification(w http.ResponseWriter, r *http.Request) {
	n := GetNotificationCtx(r)
	render.Render(w, r, NewNotificationResponse(n))
}

func ListNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := models.GetAllNotifications(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewNotificationListResponse(notifications)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateNotification(w http.ResponseWriter, r *http.Request) {
	n := GetNotificationCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	expireAt, err := http.ParseTime(data.ExpireAt)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	// Make sure expireAt is in the future
	if expireAt.Before(time.Now()) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("expire_at must be in the future")))
		return
	}

	n.CID = data.CID
	n.Category = data.Category
	n.Title = data.Title
	n.Body = data.Body
	n.ExpireAt = expireAt

	if err := n.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNotificationResponse(n))
}

func PatchNotification(w http.ResponseWriter, r *http.Request) {
	n := GetNotificationCtx(r)
	data := &Request{}
	if err := data.Bind(r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if data.CID != 0 {
		n.CID = data.CID
	}
	if data.Category != "" {
		n.Category = data.Category
	}
	if data.Title != "" {
		n.Title = data.Title
	}
	if data.Body != "" {
		n.Body = data.Body
	}
	if data.ExpireAt != "" {
		expireAt, err := http.ParseTime(data.ExpireAt)
		if err != nil {
			render.Render(w, r, utils.ErrInvalidRequest(err))
			return
		}

		// Make sure expireAt is in the future
		if expireAt.Before(time.Now()) {
			render.Render(w, r, utils.ErrInvalidRequest(errors.New("expire_at must be in the future")))
			return
		}

		n.ExpireAt = expireAt
	}

	if err := n.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewNotificationResponse(n))
}

func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	n := GetNotificationCtx(r)
	if err := n.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}
}
