package dto

import (
	"encoding/json"
	"io"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/entity"
)

type SellerRequest struct {
	Company     string  `json:"company" validate:"required"`
	Phone       *string `json:"phone" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	District    string  `json:"district" validate:"required"`
	Email       string  `json:"emali" validate:"required"`
	PICName     string  `json:"pic_name" validate:"required"`
	PICPhone    *string `json:"pic_phone" validate:"required"`
	NPWP        string  `json:"npwp" validate:"required"`
	KTP         string  `json:"ktp" validate:"required"`
	NoIUP       string  `json:"no_iup" validate:"required"`
	ValidPeriod string  `json:"valid_period" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Status      *bool   `json:"status" validate:"required"`
}

func (s *SellerRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(s)
}

func (s *SellerRequest) ToEntity() (*entity.Seller, error) {
	validPeriod, err := time.Parse("2006-01-02", s.ValidPeriod)
	if err != nil {
		return nil, err
	}

	return &entity.Seller{
		Company:     s.Company,
		Phone:       s.Phone,
		Address:     s.Address,
		District:    s.District,
		Email:       s.Email,
		PICName:     s.PICName,
		PICPhone:    s.PICPhone,
		NPWP:        s.NPWP,
		KTP:         s.KTP,
		NoIUP:       s.NoIUP,
		ValidPeriod: validPeriod,
		Description: s.Description,
		Status:      *s.Status,
	}, nil
}
