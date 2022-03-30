package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"

type AuthTokenResponse struct {
	Username      string `json:"username"`
	PrivLevel     int64  `json:"role_id"`
	PrivLevelName string `json:"role_name"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
}

func NewAuthTokenResponse(at string, rt string, user *entity.UserData) *AuthTokenResponse {
	return &AuthTokenResponse{
		Username:      user.Username,
		PrivLevel:     user.Role,
		PrivLevelName: user.RoleName,
		AccessToken:   at,
		RefreshToken:  rt,
	}
}
