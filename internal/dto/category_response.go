package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"

type CategoryResponse struct {
	Id     uint64 `json:"id"`
	Name   string `json:"category_name"`
	Pajak  int64  `json:"tax"`
	Status bool   `json:"status"`
}

type CategoriesResponse []*CategoryResponse

func NewCategoryReponse(category entity.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:     category.Id,
		Name:   category.Name,
		Pajak:  category.Pajak,
		Status: category.Status,
	}
}

func NewCategoriesResponse(categories entity.Categories) CategoriesResponse {
	data := CategoriesResponse{}

	for _, category := range categories {
		temp := NewCategoryReponse(*category)
		data = append(data, temp)
	}

	return data
}
