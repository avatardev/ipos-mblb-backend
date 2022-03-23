package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"

type CategoryResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"category_name"`
	Pajak  int64  `json:"tax"`
	Status bool   `json:"status"`
}

type CategoriesResponse struct {
	Category []*CategoryResponse `json:"category"`
	Offset   uint64              `json:"offset"`
	Limit    uint64              `json:"limit"`
	Total    uint64              `json:"total"`
}

func NewCategoryReponse(category entity.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:     category.Id,
		Name:   category.Name,
		Pajak:  category.Pajak,
		Status: category.Status,
	}
}

func NewCategoriesResponse(categories entity.Categories, limit uint64, offset uint64, total uint64) *CategoriesResponse {
	data := &CategoriesResponse{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	for _, category := range categories {
		temp := NewCategoryReponse(*category)
		data.Category = append(data.Category, temp)
	}

	return data
}
