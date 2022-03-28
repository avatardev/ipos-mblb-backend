package auth

import (
	"log"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authData := &dto.UserPostRequest{}
		if err := authData.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), authData); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}

			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := a.Service.Login(r.Context(), authData)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (a *AuthHandler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshToken := &dto.AuthRefreshToken{}
		if err := refreshToken.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), refreshToken); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}

			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := a.Service.RefreshToken(r.Context(), refreshToken)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
