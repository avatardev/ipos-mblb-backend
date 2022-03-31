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
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CASHIER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Printf("[GenerateDetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := time.Parse("2006-01-02", endDate)
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
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CASHIER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Printf("[GenerateBriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := time.Parse("2006-01-02", endDate)
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

func (o *OrderHandler) DetailTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CASHIER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offestParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[DetailTrx] error: %v\n", err)
			if offset == "" {
				offestParsed = 0
			}
		}

		res, err := o.Service.DetailTrx(r.Context(), startParsed, endParsed, limitParsed, offestParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusCreated, res)
	}
}

func (o *OrderHandler) BriefTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CASHIER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offestParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[BriefTrx] error: %v\n", err)
			if offset == "" {
				offestParsed = 0
			}
		}

		res, err := o.Service.BriefTrx(r.Context(), startParsed, endParsed, limitParsed, offestParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (o *OrderHandler) InsertNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CASHIER) {
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
