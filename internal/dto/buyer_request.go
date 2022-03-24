package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/entity"
)

type BuyerRequest struct {
	VehiclePlate string `json:"vehicle_plate" validate:"required"`
	CategoryId   uint64 `json:"category_id" validate:"required"`
	Company      string `json:"company" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PicName      string `json:"pic_name" validate:"required"`
	PicPhone     string `json:"pic_phone" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Status       *bool  `json:"status" validate:"required"`
}

func (b *BuyerRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(b)
}

func (b *BuyerRequest) ToEntity() entity.Buyer {
	return entity.Buyer{
		VehiclePlate:      b.VehiclePlate,
		VehicleCategoryId: b.CategoryId,
		Company:           b.Company,
		Phone:             b.Phone,
		Address:           b.Address,
		Email:             b.Email,
		PICName:           b.PicName,
		PICPhone:          b.PicPhone,
		Description:       b.Description,
		Status:            *b.Status,
	}
}
