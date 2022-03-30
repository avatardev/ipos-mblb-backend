package entity

type User struct {
	Id       int64
	Username string
	Password string
	VPlate   *string
	SellerId *int64
}

type Users []*User
