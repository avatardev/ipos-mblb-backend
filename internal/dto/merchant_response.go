package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/merchant/entity"

type MerchantResponse struct {
	Id          int64   `json:"id"`
	ProductId   int64   `json:"product_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
}

type MerchantsResponse struct {
	MerchantItem []*MerchantResponse `json:"merchant_item"`
	Offset       uint64              `json:"offset"`
	Limit        uint64              `json:"limit"`
	Total        uint64              `json:"total"`
}

func NewMerchantResponse(item *entity.MerchantItem) *MerchantResponse {
	return &MerchantResponse{
		Id:          item.Id,
		ProductId:   item.ProductId,
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		Status:      item.Status,
	}
}

func NewMerchantsResponse(items entity.MerchantItems, limit uint64, offset uint64, total uint64) *MerchantsResponse {
	data := &MerchantsResponse{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	for _, itm := range items {
		temp := NewMerchantResponse(itm)
		data.MerchantItem = append(data.MerchantItem, temp)
	}

	return data
}
