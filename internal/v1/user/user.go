package user

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
	CID              uint   `json:"cid" example:"1293257" validate:"required"`
	FirstName        string `json:"first_name" example:"Raaj" validate:"required"`
	LastName         string `json:"last_name" example:"Patel" validate:"required"`
	PreferredName    string `json:"preferred_name" example:"Raaj"`
	Email            string `json:"email" example:"vatusa6@vatusa.net" validate:"required,email"`
	PreferredOIs     string `json:"preferred_ois" example:"RP"`
	PilotRating      uint   `json:"pilot_rating" example:"1" validate:"required"`
	ControllerRating uint   `json:"controller_rating" example:"1" validate:"required"`
	DiscordID        string `json:"discord_id" example:"1234567890"`
	LastCertSync     string `json:"last_cert_sync" example:"2021-01-01T00:00:00Z"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r *http.Request) error {
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
	if res.User == nil {
		return errors.New("missing required user")
	}
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
	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	user := &models.User{
		CID:              req.CID,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		PreferredName:    req.PreferredName,
		Email:            req.Email,
		PreferredOIs:     req.PreferredOIs,
		PilotRating:      req.PilotRating,
		ControllerRating: req.ControllerRating,
		DiscordID:        req.DiscordID,
	}
	if err := user.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewUserResponse(user))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserCtx(r)

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
	user := GetUserCtx(r)

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	user.CID = req.CID
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PreferredName = req.PreferredName
	user.Email = req.Email
	user.PreferredOIs = req.PreferredOIs
	user.PilotRating = req.PilotRating
	user.ControllerRating = req.ControllerRating
	user.DiscordID = req.DiscordID

	if err := user.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewUserResponse(user))
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserCtx(r)

	req := &Request{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.PreferredName != "" {
		user.PreferredName = req.PreferredName
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.PreferredOIs != "" {
		user.PreferredOIs = req.PreferredOIs
	}
	if req.DiscordID != "" {
		user.DiscordID = req.DiscordID
	}

	if err := user.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Render(w, r, NewUserResponse(user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserCtx(r)

	if err := user.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInternalServer)
		return
	}

	render.Status(r, http.StatusNoContent)
}
