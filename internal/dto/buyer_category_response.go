package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/buyer_category/entity"

type BuyerCategoryResponse struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	IsMultiProduct bool   `json:"multi_product"`
}

type BuyerCategoriesResponse struct {
	Category []*BuyerCategoryResponse `json:"category"`
	Offset   uint64                   `json:"offset"`
	Limit    uint64                   `json:"limit"`
	Total    uint64                   `json:"total"`
}

func NewBuyerCategoryResponse(category *entity.BuyerCategory) *BuyerCategoryResponse {
	return &BuyerCategoryResponse{
		Id:             category.Id,
		Name:           category.Name,
		IsMultiProduct: category.IsMultiProduct,
	}
}

func NewBuyerCategoriesResponse(categories entity.BuyersCategories, limit uint64, offset uint64, total uint64) *BuyerCategoriesResponse {
	res := &BuyerCategoriesResponse{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	for _, category := range categories {
		temp := NewBuyerCategoryResponse(category)
		res.Category = append(res.Category, temp)
	}

	return res
}
