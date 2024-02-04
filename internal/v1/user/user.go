package user

import (
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"net/http"
)

type Request struct {
	*models.User
}

func (req *Request) Bind(r *http.Request) error {
	if req.User == nil {
		return errors.New("missing required User fields")
	}

	return nil
}

type Response struct {
	*models.User
}

func NewUserResponse(user *models.User) *Response {
	resp := &Response{User: user}

	return resp
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserListResponse(users []models.User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, NewUserResponse(&user))
	}
	return list
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	user := data.User
	if err := user.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewUserResponse(user))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	render.Render(w, r, NewUserResponse(user))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	if err := render.RenderList(w, r, NewUserListResponse(users)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	data := &Request{User: user}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}
	user = data.User
	if err := user.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewUserResponse(user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	if err := user.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.Render(w, r, nil)
}
