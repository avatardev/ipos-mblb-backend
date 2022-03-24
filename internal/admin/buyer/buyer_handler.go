package buyer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/go-playground/validator/v10"
)

type BuyerHandler struct {
	Service BuyerService
}

func (b *BuyerHandler) GetBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetBuyer] error: %v\n", err)
			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetBuyer] error: %v\n", err)
			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")

		res, err := b.Service.GetBuyer(r.Context(), keyword, limitParsed, offsetParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidResources)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) StoreBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buyer := &dto.BuyerRequest{}
		if err := buyer.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), buyer); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := b.Service.StoreBuyer(r.Context(), buyer)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) PingBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := b.Service.Ping(r.Context())
		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) PingError() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := b.Service.PingError(r.Context())
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
		}
	}
}

func NewBuyerHandler(service BuyerService) *BuyerHandler {
	return &BuyerHandler{Service: service}
}
