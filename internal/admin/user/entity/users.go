package entity

type User struct {
	Id       int64
	Username string
	Password string
}

type Users []*User
