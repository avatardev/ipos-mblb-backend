package dto

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthTokenResponse(at string, rt string) *AuthTokenResponse {
	return &AuthTokenResponse{
		AccessToken:  at,
		RefreshToken: rt,
	}
}
