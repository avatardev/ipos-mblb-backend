package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"
)

type CategoryRequest struct {
	Name   string `json:"category_name" validate:"required"`
	Pajak  int64  `json:"tax" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

func (c *CategoryRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}

func (c *CategoryRequest) ToEntity() entity.Category {
	return entity.Category{
		Name:   c.Name,
		Pajak:  c.Pajak,
		Status: c.Status,
	}
}
