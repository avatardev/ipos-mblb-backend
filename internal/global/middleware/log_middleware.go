package middleware

import (
	"context"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log"
	logEntity "github.com/avatardev/ipos-mblb-backend/internal/admin/activity_log/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/gorilla/mux"
)

func ActivityLogMiddleware(service activity_log.LogService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logInfo := &dto.LogInfo{}
			var key logEntity.LogCtxKey = "log-info"

			// assign context with log-related metadata
			ctx := context.WithValue(r.Context(), key, logInfo)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			info, ok := r.Context().Value(logEntity.LogCtxKey("log-info")).(*dto.LogInfo)
			if !ok || info.Message == "" {
				// don't log anything on non-successful response
				return
			}

			service.Store(r.Context(), info)
		})
	}
}
