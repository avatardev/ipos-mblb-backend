package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"
)

type ProductRequest struct {
	CategoryId  int64   `json:"category_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Price       float32 `json:"price_m3" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Status      *bool   `json:"status" validate:"required"`
}

func (p *ProductRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *ProductRequest) ToEntity() entity.Product {
	return entity.Product{
		CategoryId:  p.CategoryId,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		Status:      *p.Status,
	}
}
