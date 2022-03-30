package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/user/entity"
)

type UserPostRequest struct {
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password" validate:"required"`
	VPlate   *string `json:"vehicle_plate"`
	SellerId *int64  `json:"seller_id"`
}

func (ur *UserPostRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ur)
}

func (ur *UserPostRequest) ToEntity() entity.User {
	res := entity.User{
		Username: ur.Username,
		Password: ur.Password,
	}

	if ur.SellerId != nil {
		res.SellerId = ur.SellerId
	} else if ur.VPlate != nil {
		res.VPlate = ur.VPlate
	}

	return res
}

type UserPutRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}

func (ur *UserPutRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ur)
}

func (ur *UserPutRequest) ToEntity() entity.User {
	res := entity.User{
		Username: ur.Username,
	}

	if ur.Password != "" {
		res.Password = ur.Password
	}

	return res
}
