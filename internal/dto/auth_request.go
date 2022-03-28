package dto

import (
	"encoding/json"
	"io"
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
