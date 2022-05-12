package order

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/timeutil"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	Service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (o *OrderHandler) GenerateDetailTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
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

		res, err := o.Service.GenerateDetailTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		filename := fmt.Sprintf("ExportDetailPeriode%v-%v", startDate, endDate)
		responseutil.WriteFileResponse(w, http.StatusCreated, filename, *res)
	}
}

func (o *OrderHandler) GenerateBriefTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateBriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateBriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.GenerateBriefTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		filename := fmt.Sprintf("ExportPeriode%v-%v", startDate, endDate)
		responseutil.WriteFileResponse(w, http.StatusCreated, filename, *res)
	}
}

func (o *OrderHandler) GenerateMonitorTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateMonitorTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[GenerateMonitorTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.GenerateMonitorTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		filename := fmt.Sprintf("ExportPembandingPeriode%v-%v", startDate, endDate)
		responseutil.WriteFileResponse(w, http.StatusCreated, filename, *res)
	}
}

func (o *OrderHandler) GenerateDailyTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		sellerId, ok := mux.Vars(r)["sellerId"]
		if !ok {
			log.Print("[GenerateDailyTrx] error: seller id not found\n")
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(sellerId, 10, 64)
		if err != nil {
			log.Printf("[GenerateDailyTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := o.Service.GenerateDailyTrx(r.Context(), parsedId)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		filename := fmt.Sprintf("ExportLaporanHarian%v", time.Now().Format("02-01-2006"))
		responseutil.WriteFileResponse(w, http.StatusCreated, filename, *res)
	}
}

func (o *OrderHandler) DetailTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.DetailTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusCreated, res)
	}
}

func (o *OrderHandler) BriefTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.BriefTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (o *OrderHandler) MonitorTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 00:00:00", startDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[MonitorTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := timeutil.ParseLocalTime(fmt.Sprintf("%s 23:59:59", endDate), "2006-01-02 15:04:05")
		if err != nil {
			log.Printf("[MonitorTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.MonitorTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (o *OrderHandler) DailyTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		sellerId, ok := mux.Vars(r)["sellerId"]
		if !ok {
			log.Print("[DailyTrx] error: seller id not found\n")
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(sellerId, 10, 64)
		if err != nil {
			log.Printf("[DailyTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := o.Service.DailyTrx(r.Context(), parsedId)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (o *OrderHandler) InsertNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_SELLER, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		id, exist := mux.Vars(r)["orderId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		req := &dto.InsertNoteRequest{}
		if err := req.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), req); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.InsertNote(r.Context(), parsedId, req.Note)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
