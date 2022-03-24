package dto

import (
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/entity"
)

type BuyerResponse struct {
	VehiclePlate string `json:"vehicle_plate"`
	Category     string `json:"category"`
	Company      string `json:"company"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	PicName      string `json:"pic_name"`
	PicPhone     string `json:"pic_phone"`
	Description  string `json:"description"`
	Status       bool   `json:"status"`
}

type BuyersResponse struct {
	Buyer  []*BuyerResponse `json:"buyer"`
	Offset uint64           `json:"offset"`
	Limit  uint64           `json:"limit"`
	Total  uint64           `json:"total"`
}

func NewBuyerResponse(buyer entity.Buyer) *BuyerResponse {
	return &BuyerResponse{
		VehiclePlate: buyer.VehiclePlate,
		Category:     buyer.VehicleCategoryName,
		Company:      buyer.Company,
		Phone:        buyer.Phone,
		Address:      buyer.Address,
		Email:        buyer.Email,
		PicName:      buyer.PICName,
		PicPhone:     buyer.PICPhone,
		Description:  buyer.Description,
		Status:       buyer.Status,
	}
}

func NewBuyersResponse(buyers entity.Buyers, count uint64, limit uint64, offset uint64) *BuyersResponse {
	res := &BuyersResponse{
		Offset: offset,
		Limit:  limit,
		Total:  count,
	}

	for _, buyer := range buyers {
		temp := NewBuyerResponse(*buyer)
		res.Buyer = append(res.Buyer, temp)
	}
	log.Println(len(buyers))
	return res
}
