package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/entity"
)

type MerchantRequest struct {
	ProductId   int64   `json:"product_id" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description"`
	Status      *bool   `json:"status" validate:"required"`
}

func (m *MerchantRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(m)
}

func (m *MerchantRequest) ToEntity() entity.MerchantItem {
	return entity.MerchantItem{
		ProductId:   m.ProductId,
		Price:       m.Price,
		Description: m.Description,
		Status:      *m.Status,
	}
}
