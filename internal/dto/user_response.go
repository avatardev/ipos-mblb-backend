package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/user/entity"

type UserResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type UsersResponse struct {
	Users  []*UserResponse `json:"user"`
	Limit  uint64          `json:"limit"`
	OFfset uint64          `json:"offset"`
	Total  uint64          `json:"total"`
}

func NewUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
}

func NewUsersResponse(users entity.Users, limit uint64, offset uint64, total uint64) *UsersResponse {
	res := &UsersResponse{
		Limit:  limit,
		OFfset: offset,
		Total:  total,
	}

	for _, user := range users {
		res.Users = append(res.Users, NewUserResponse(user))
	}

	return res
}
