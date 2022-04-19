package impl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/internal/global/config"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/logutil"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	Ar AuthRepositoryImpl
}

func (a *AuthServiceImpl) Login(ctx context.Context, req *dto.UserPostRequest) (*dto.AuthTokenResponse, error) {
	user := req.ToEntity()

	userData, err := a.Ar.GetByUsername(ctx, user.Username)
	if err != nil || userData == nil {
		return nil, errors.ErrUserCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		log.Printf("[Auth.Login] error: %v\n", err)
		return nil, errors.ErrUserCredential
	}

	accessToken, err := a.newAccessToken(userData)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.newRefreshToken(userData)
	if err != nil {
		return nil, err
	}

	logutil.GenerateActivityLogNoAuth(ctx, userData.Id, "login to system")
	return dto.NewAuthTokenResponse(accessToken, refreshToken, userData), nil
}

func (a *AuthServiceImpl) RefreshToken(ctx context.Context, req *dto.AuthRefreshToken) (*dto.AuthTokenResponse, error) {
	config := config.GetConfig()
	refreshToken := req.Token

	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != config.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signin method invalid")
		}
		return config.JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		log.Printf("[Auth.RefreshToken] error: %v\n", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Printf("[Auth.RefreshToken] error: %v\n", err)
		return nil, err
	}

	data := claims["data"].(map[string]interface{})
	userId := int64(data["id"].(float64))

	user, err := a.Ar.GetById(ctx, userId)
	if err != nil {
		log.Printf("[Auth.RefreshToken] error: %v\n", err)
		return nil, errors.ErrUserCredential
	}

	at, err := a.newAccessToken(user)
	if err != nil {
		return nil, err
	}

	rt, err := a.newRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return dto.NewAuthTokenResponse(at, rt, user), nil
}

func (a *AuthServiceImpl) FindUserByAccessToken(ctx context.Context, accessToken string) (*dto.AuthUserLevel, error) {
	config := config.GetConfig()

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != config.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signin method invalid")
		}
		return config.JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		log.Printf("[Auth.FindUserByAccessToken] error: %v\n", err)
		return nil, errors.ErrTokenExpired
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Printf("[Auth.FindUserByAccessToken] error: %v\n", err)
		return nil, err
	}

	data := claims["data"].(map[string]interface{})
	userId := int64(data["id"].(float64))

	user, err := a.Ar.GetById(ctx, userId)
	if err != nil {
		log.Printf("[Auth.FindUserByAccessToken] error: %v\n", err)
		return nil, errors.ErrUserCredential
	}

	return dto.NewAuthUserLevel(user), nil
}

func (a *AuthServiceImpl) newAccessToken(user *entity.UserData) (string, error) {
	config := config.GetConfig()

	claims := a.newUserClaim(user.Id, user.Username, config.JWT_AT_EXPIRATION)
	accessToken := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)
	signed, err := accessToken.SignedString(config.JWT_SIGNATURE_KEY)
	if err != nil {
		log.Printf("[Auth.NewAccessToken] error: %v\n", err)
		return "", err
	}

	return signed, nil
}

func (a *AuthServiceImpl) newRefreshToken(user *entity.UserData) (string, error) {
	config := config.GetConfig()

	claims := a.newUserClaim(user.Id, user.Username, config.JWT_RT_EXPIRATION)
	refreshToken := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)
	signed, err := refreshToken.SignedString(config.JWT_SIGNATURE_KEY)
	if err != nil {
		log.Printf("[Auth.NewRefreshToken] error: %v\n", err)
		return "", err
	}

	return signed, nil
}

func (a *AuthServiceImpl) newUserClaim(id int64, username string, exp time.Duration) *jwt.MapClaims {
	return &jwt.MapClaims{
		"iss": config.GetConfig().JWT_ISSUER,
		"exp": time.Now().Add(exp).Unix(),
		"data": map[string]interface{}{
			"id":       id,
			"username": username,
		},
	}
}
