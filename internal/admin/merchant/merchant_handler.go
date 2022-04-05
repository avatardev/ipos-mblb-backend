package merchant

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

type MerchantHandler struct {
	Service MerchantService
}

func NewMerchantHandler(service MerchantService) *MerchantHandler {
	return &MerchantHandler{Service: service}
}

func (m *MerchantHandler) GetMerchant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")
		sellerId, exist := mux.Vars(r)["sellerId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
		}

		sellerParsed, err := strconv.ParseInt(sellerId, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
		}

		res, err := m.Service.GetMerchant(r.Context(), sellerParsed, keyword, limitParsed, offsetParsed)
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

func (m *MerchantHandler) GetMerchantById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		sellerId, exist := mux.Vars(r)["sellerId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
		}

		sellerParsed, err := strconv.ParseInt(sellerId, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
		}

		itemId, exist := mux.Vars(r)["itemId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
		}

		itemParsed, err := strconv.ParseInt(itemId, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
		}

		res, err := m.Service.GetMerchantById(r.Context(), sellerParsed, itemParsed)
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

func (m *MerchantHandler) UpdateMerchant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		sellerId, exist := mux.Vars(r)["sellerId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
		}

		sellerParsed, err := strconv.ParseInt(sellerId, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
		}

		itemId, exist := mux.Vars(r)["itemId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
		}

		itemParsed, err := strconv.ParseInt(itemId, 10, 64)
		if err != nil {
			log.Printf("[GetMerchant] error: %v\n", err)
		}

		item := &dto.MerchantRequest{}
		if err := item.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), item); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := m.Service.UpdateMerchant(r.Context(), sellerParsed, itemParsed, item)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
