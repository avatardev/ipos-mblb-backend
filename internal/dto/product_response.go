package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"

type ProductResponse struct {
	Id          int64   `json:"id"`
	CategoryId  int64   `json:"category_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price_m3"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
}

type ProductsResponse []*ProductResponse

func NewProductResponse(product entity.Product) *ProductResponse {
	return &ProductResponse{
		Id:          product.Id,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Description: product.Description,
		Status:      product.Status,
	}
}

func NewProductsResponse(products entity.Products) ProductsResponse {
	data := ProductsResponse{}

	for _, product := range products {
		temp := NewProductResponse(*product)
		data = append(data, temp)
	}

	return data
}
