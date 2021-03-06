package dto

import (
	"github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"
)

type ProductResponse struct {
	Id           int64   `json:"id"`
	CategoryId   int64   `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Name         string  `json:"name"`
	Price        float32 `json:"price_m3"`
	Tax          uint64  `json:"tax"`
	Description  string  `json:"description"`
	Status       bool    `json:"status"`
	Img          *string `json:"img"`
}

type ProductsResponse struct {
	Products []*ProductResponse `json:"product"`
	Offset   uint64             `json:"offset"`
	Limit    uint64             `json:"limit"`
	Total    uint64             `json:"total"`
}

func NewProductResponse(product entity.Product) *ProductResponse {
	return &ProductResponse{
		Id:           product.Id,
		CategoryId:   product.CategoryId,
		CategoryName: product.CategoryName,
		Name:         product.Name,
		Price:        product.Price,
		Tax:          product.Tax,
		Description:  product.Description,
		Status:       product.Status,
		Img:          product.Img,
	}
}

func NewProductsResponse(products entity.Products, limit uint64, offset uint64, total uint64) *ProductsResponse {
	data := &ProductsResponse{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	for _, product := range products {
		temp := NewProductResponse(*product)
		data.Products = append(data.Products, temp)
	}

	return data
}
