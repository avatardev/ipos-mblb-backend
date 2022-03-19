package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/category/entity"

type Category struct {
	Id     uint64 `json:"id"`
	Name   string `json:"category_name"`
	Pajak  int64  `json:"tax"`
	Status bool   `json:"status"`
}

type Categories []*Category

func MapToCategory(category entity.Category) *Category {
	return &Category{
		Id:     category.Id,
		Name:   category.Name,
		Pajak:  category.Pajak,
		Status: category.Status,
	}
}

func MapToCategories(categories entity.Categories) Categories {
	data := Categories{}

	for _, category := range categories {
		temp := MapToCategory(*category)
		data = append(data, temp)
	}

	return data
}
