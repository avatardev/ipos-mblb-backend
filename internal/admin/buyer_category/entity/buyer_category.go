package entity

type BuyerCategory struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	IsMultiProduct bool   `json:"multi_product"`
}

type BuyersCategories []*BuyerCategory
