package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
)

type AuthUserLevel struct {
	Id       int64
	Username string
	Role     int64
}

type AuthRefreshToken struct {
	Token string `json:"refresh_token" validate:"required"`
}

func (art *AuthRefreshToken) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(art)
}

func NewAuthUserLevel(user *entity.UserData) *AuthUserLevel {
	return &AuthUserLevel{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}
}
