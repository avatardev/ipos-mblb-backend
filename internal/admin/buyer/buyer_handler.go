package buyer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type BuyerHandler struct {
	Service BuyerService
}

func NewBuyerHandler(service BuyerService) *BuyerHandler {
	return &BuyerHandler{Service: service}
}

func (b *BuyerHandler) GetBuyerName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		res, err := b.Service.GetBuyerName(r.Context())
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

func (b *BuyerHandler) GetBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER, privutil.USER_SELLER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

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

func (b *BuyerHandler) GetBuyerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER, privutil.USER_SELLER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		plate, exists := mux.Vars(r)["buyerId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := b.Service.GetBuyerById(r.Context(), plate)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrNotFound)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) StoreBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER, privutil.USER_SELLER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

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

func (b *BuyerHandler) UpdateBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER, privutil.USER_SELLER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		plateNumber, exists := mux.Vars(r)["buyerId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		buyer := &dto.BuyerRequest{}
		buyer.FromJSON(r.Body)

		v := validator.New()
		if err := v.StructCtx(r.Context(), buyer); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}

			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := b.Service.UpdateBuyer(r.Context(), plateNumber, buyer)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) DeleteBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER, privutil.USER_SELLER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		plate, exists := mux.Vars(r)["buyerId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		if err := b.Service.DeleteBuyer(r.Context(), plate); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "buyer deleted")
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
