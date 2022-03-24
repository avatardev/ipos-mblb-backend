package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/entity"
)

type ProductCategoryRequest struct {
	Name   string `json:"category_name" validate:"required"`
	Pajak  int64  `json:"tax" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

func (c *ProductCategoryRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}

func (c *ProductCategoryRequest) ToEntity() entity.ProductCategory {
	return entity.ProductCategory{
		Name:   c.Name,
		Pajak:  c.Pajak,
		Status: c.Status,
	}
}
