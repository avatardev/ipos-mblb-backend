package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/product/entity"

type ProductResponse struct {
	Id           int64   `json:"id"`
	CategoryName string  `json:"category_name"`
	Name         string  `json:"name"`
	Price        float32 `json:"price_m3"`
	Description  string  `json:"description"`
	Status       bool    `json:"status"`
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
		CategoryName: product.CategoryName,
		Name:         product.Name,
		Price:        product.Price,
		Description:  product.Description,
		Status:       product.Status,
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
