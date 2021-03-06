package dto

import (
	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer/entity"
)

type BuyerResponse struct {
	VehiclePlate string  `json:"vehicle_plate"`
	CategoryId   uint64  `json:"category_id"`
	Category     string  `json:"category"`
	Company      string  `json:"company"`
	Phone        *string `json:"phone"`
	Address      string  `json:"address"`
	Email        string  `json:"email"`
	PicName      string  `json:"pic_name"`
	PicPhone     *string `json:"pic_phone"`
	Description  string  `json:"description"`
	Status       bool    `json:"status"`
}

type BuyersResponse struct {
	Buyer  []*BuyerResponse `json:"buyer"`
	Offset uint64           `json:"offset"`
	Limit  uint64           `json:"limit"`
	Total  uint64           `json:"total"`
}

type BuyerCompanyResponse struct {
	VPlate  string `json:"vehicle_plate"`
	Company string `json:"company_name"`
}

type BuyersCompanyResponse struct {
	Company []*BuyerCompanyResponse `json:"company"`
}

func NewBuyerResponse(buyer entity.Buyer) *BuyerResponse {
	return &BuyerResponse{
		VehiclePlate: buyer.VehiclePlate,
		CategoryId:   buyer.VehicleCategoryId,
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
	return res
}

func NewBuyerCompanyResponse(bc *entity.BuyerCompany) *BuyerCompanyResponse {
	return &BuyerCompanyResponse{
		VPlate:  bc.VPlate,
		Company: bc.Company,
	}
}

func NewBuyersCompanyResponse(bc entity.BuyersCompany) *BuyersCompanyResponse {
	res := &BuyersCompanyResponse{}

	for _, buyer := range bc {
		temp := NewBuyerCompanyResponse(buyer)
		res.Company = append(res.Company, temp)
	}

	return res
}
