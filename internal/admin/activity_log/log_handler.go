package activity_log

import (
	"log"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/go-playground/validator/v10"
)

type LogHandler struct {
	Service LogService
}

func NewLogHandler(service LogService) *LogHandler {
	return &LogHandler{Service: service}
}

func (l *LogHandler) GetLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		res, err := l.Service.GetLogs(r.Context())
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidResources)
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (l *LogHandler) StoreLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInfo := &dto.LogInfo{}

		if err := logInfo.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON} error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), logInfo); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		if err := l.Service.Store(r.Context(), logInfo); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		responseutil.WriteSuccessResponse(w, http.StatusOK, nil)
	}
}
