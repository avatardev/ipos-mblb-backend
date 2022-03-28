package auth

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.UserPostRequest) (*dto.AuthTokenResponse, error)
	RefreshToken(ctx context.Context, req *dto.AuthRefreshToken) (*dto.AuthTokenResponse, error)
}

func NewAuthService(Ar impl.AuthRepositoryImpl) AuthService {
	return &impl.AuthServiceImpl{Ar: Ar}
}
