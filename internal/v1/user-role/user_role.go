package user_role

import (
	"encoding/json"
	"errors"
	"github.com/VATUSA/primary-api/pkg/database"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"github.com/VATUSA/primary-api/pkg/utils"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Request struct {
	CID        uint   `json:"cid" example:"1293257" validate:"required"`
	RoleID     string `json:"role_id" example:"ATM" validate:"required"`
	FacilityID string `json:"facility_id" example:"ZDV" validate:"required"`
}

func (req *Request) Validate() error {
	return validator.New().Struct(req)
}

func (req *Request) Bind(r http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		return err
	}

	return nil
}

type Response struct {
	*models.UserRole
}

func NewUserRoleResponse(r *models.UserRole) *Response {
	return &Response{UserRole: r}
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.UserRole == nil {
		return errors.New("user role not found")
	}

	return nil
}

func NewUserRoleListResponse(r []models.UserRole) []render.Renderer {
	list := []render.Renderer{}
	for _, d := range r {
		list = append(list, NewUserRoleResponse(&d))
	}
	return list
}

func CreateUserRoles(w http.ResponseWriter, r *http.Request) {
	req := &Request{}
	if err := req.Bind(*r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if !models.IsValidRole(database.DB, req.RoleID) {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid role")))
		return
	}

	userRole := &models.UserRole{
		CID:        req.CID,
		RoleID:     req.RoleID,
		FacilityID: req.FacilityID,
	}

	if err := userRole.Create(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewUserRoleResponse(userRole))
}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	userRole := GetUserRoleCtx(r)

	render.Render(w, r, NewUserRoleResponse(userRole))
}

func ListUserRoles(w http.ResponseWriter, r *http.Request) {
	userRoles, err := models.GetAllUserRoles(database.DB)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := render.RenderList(w, r, NewUserRoleListResponse(userRoles)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	userRole := GetUserRoleCtx(r)

	req := &Request{}
	if err := req.Bind(*r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	userRole.CID = req.CID
	userRole.RoleID = req.RoleID
	userRole.FacilityID = req.FacilityID

	if err := userRole.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewUserRoleResponse(userRole))
}

func PatchUserRole(w http.ResponseWriter, r *http.Request) {
	userRole := GetUserRoleCtx(r)

	req := &Request{}
	if err := req.Bind(*r); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if req.CID != 0 {
		userRole.CID = req.CID
	}
	if req.RoleID != "" {
		if !models.IsValidRole(database.DB, req.RoleID) {
			render.Render(w, r, utils.ErrInvalidRequest(errors.New("invalid role")))
			return
		}
		userRole.RoleID = req.RoleID
	}
	if req.FacilityID != "" {
		userRole.FacilityID = req.FacilityID
	}

	if err := userRole.Update(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewUserRoleResponse(userRole))
}

func DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	userRole := GetUserRoleCtx(r)

	if err := userRole.Delete(database.DB); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
