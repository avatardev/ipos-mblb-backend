package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/entity"
)

type BuyerCategoryRequest struct {
	Name           string `json:"category_name" validate:"required"`
	IsMultiProduct *bool  `json:"multi_product" validate:"required"`
}

func (bc *BuyerCategoryRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(bc)
}

func (bc *BuyerCategoryRequest) ToEntity() entity.BuyerCategory {
	return entity.BuyerCategory{
		Name:           bc.Name,
		IsMultiProduct: *bc.IsMultiProduct,
	}
}
