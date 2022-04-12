package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/gorilla/mux"
)

func AuthMiddleware(service auth.AuthService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[AccessControl-Middleware] Incoming Request to %v\n", r.URL)

			token := r.Header.Get("Authorization")
			if token == "" {
				responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
				return
			}

			splittedToken := strings.Split(token, " ")
			if len(splittedToken) != 2 {
				responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
				return
			}

			if splittedToken[0] != "Bearer" {
				responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
				return
			}

			accessToken := splittedToken[1]
			user, err := service.FindUserByAccessToken(r.Context(), accessToken)
			if err != nil {
				responseutil.WriteErrorResponse(w, err)
				return
			} else if user == nil {
				responseutil.WriteErrorResponse(w, errors.ErrUserCredential)
				return
			}

			log.Printf("[AccessControl-Middleware] user logged in: %v (role: %v)\n", user.Username, user.Role)

			var auth entity.AuthLevelCtxKey = "user-auth"
			newCtx := context.WithValue(r.Context(), auth, user)
			r = r.WithContext(newCtx)
			next.ServeHTTP(w, r)
		})
	}
}
