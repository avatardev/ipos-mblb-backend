package dto

import (
	"github.com/avatardev/ipos-mblb-backend/internal/admin/seller/entity"
)

type SellerResponse struct {
	Id          int64   `json:"id"`
	Company     string  `json:"company"`
	Phone       *string `json:"phone"`
	Address     string  `json:"address"`
	District    string  `json:"district"`
	Email       string  `json:"emali"`
	PICName     string  `json:"pic_name"`
	PICPhone    *string `json:"pic_phone"`
	NPWP        string  `json:"npwp"`
	KTP         string  `json:"ktp"`
	NoIUP       string  `json:"no_iup"`
	ValidPeriod string  `json:"valid_period"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
}

type SellersResponse struct {
	Seller []*SellerResponse `json:"seller"`
	Offset uint64            `json:"offset"`
	Limit  uint64            `jsoN:"limit"`
	Total  uint64            `json:"total"`
}

type SellerCompanyResponse struct {
	Id      int64  `json:"seller_id"`
	Company string `json:"company_name"`
}
type SellersCompanyResponse struct {
	Company []*SellerCompanyResponse `json:"company"`
}

func NewSellerResponse(seller *entity.Seller) *SellerResponse {
	return &SellerResponse{
		Id:          seller.Id,
		Company:     seller.Company,
		Phone:       seller.Phone,
		Address:     seller.Address,
		District:    seller.District,
		Email:       seller.Email,
		PICName:     seller.PICName,
		PICPhone:    seller.PICPhone,
		NPWP:        seller.NPWP,
		KTP:         seller.KTP,
		NoIUP:       seller.NoIUP,
		ValidPeriod: seller.ValidPeriod.Format("2006-01-02"),
		Description: seller.Description,
		Status:      seller.Status,
	}
}

func NewSellersResponse(sellers entity.Sellers, limit uint64, offset uint64, total uint64) *SellersResponse {
	res := &SellersResponse{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	for _, seller := range sellers {
		temp := NewSellerResponse(seller)
		res.Seller = append(res.Seller, temp)
	}

	return res
}

func NewSellerCompanyResponse(sc *entity.CompanySeller) *SellerCompanyResponse {
	return &SellerCompanyResponse{
		Id:      sc.Id,
		Company: sc.Company,
	}
}

func NewSellersCompanyResponse(sc entity.CompanySellers) *SellersCompanyResponse {
	res := &SellersCompanyResponse{}

	for _, seller := range sc {
		temp := NewSellerCompanyResponse(seller)
		res.Company = append(res.Company, temp)
	}

	return res
}
