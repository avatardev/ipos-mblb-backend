package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/product_category/entity"

type ProductCategoryResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"category_name"`
	Pajak  int64  `json:"tax"`
	Status bool   `json:"status"`
}

type ProductCategoriesResponse struct {
	Category []*ProductCategoryResponse `json:"category"`
	Offset   uint64                     `json:"offset"`
	Limit    uint64                     `json:"limit"`
	Total    uint64                     `json:"total"`
}

func NewProductCategoryResponse(category entity.ProductCategory) *ProductCategoryResponse {
	return &ProductCategoryResponse{
		Id:     category.Id,
		Name:   category.Name,
		Pajak:  category.Pajak,
		Status: category.Status,
	}
}

func NewProductCategoriesResponse(categories entity.ProductCategories, limit uint64, offset uint64, total uint64) *ProductCategoriesResponse {
	data := &ProductCategoriesResponse{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	for _, category := range categories {
		temp := NewProductCategoryResponse(*category)
		data.Category = append(data.Category, temp)
	}

	return data
}
