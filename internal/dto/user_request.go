package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/user/entity"
)

type UserPostRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (ur *UserPostRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ur)
}

func (ur *UserPostRequest) ToEntity() entity.User {
	return entity.User{
		Username: ur.Username,
		Password: ur.Password,
	}
}

type UserPutRequest struct {
	Username string `json:"username" validate:"required"`
}

func (ur *UserPutRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ur)
}

func (ur *UserPutRequest) ToEntity() entity.User {
	return entity.User{
		Username: ur.Username,
	}
}
