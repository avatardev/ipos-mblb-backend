package entity

type MerchantItem struct {
	Id          int64
	ProductId   int64
	Name        string
	Price       float32
	Description string
	Status      bool
}

type MerchantItems []*MerchantItem
