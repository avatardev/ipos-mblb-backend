package activity_log

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/timeutil"
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

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateDetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateDetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := l.Service.GetLogs(r.Context(), startParsed, endParsed)
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
