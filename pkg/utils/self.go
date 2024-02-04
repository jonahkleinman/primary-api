package utils

import (
	"encoding/json"
	"github.com/VATUSA/primary-api/pkg/database/models"
	"net/http"
)

func IsGuest(r *http.Request) bool {
	return r.Header.Get("x-guest") == "true"
}

func GetSelf(r *http.Request) *models.User {
	// Get x-user from header and cast to models.User
	if IsGuest(r) {
		return nil
	}

	user := r.Header.Get("x-user")
	if user == "" {
		return nil
	}

	xuser := models.User{}
	if err := json.Unmarshal([]byte(user), &xuser); err != nil {
		return nil
	}

	return &xuser
}

func GetSelfCID(r *http.Request) uint {
	return GetSelf(r).CID
}
