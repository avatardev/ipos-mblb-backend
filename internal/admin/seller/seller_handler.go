package seller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type SellerHandler struct {
	Service SellerService
}

func NewSellerHandler(service SellerService) *SellerHandler {
	return &SellerHandler{Service: service}
}

func (s *SellerHandler) GetSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetSeller] limit: %v, error:%v\n", limit, err)

			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetSeller] offset: %v, error: %v\n", offset, err)

			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")
		res, err := s.Service.GetSeller(r.Context(), keyword, limitParsed, offsetParsed)
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

func (s *SellerHandler) GetSellerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exists := mux.Vars(r)["sellerId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := s.Service.GetSellerById(r.Context(), parsedId)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (s *SellerHandler) StoreSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		seller := &dto.SellerRequest{}
		if err := seller.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), seller); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := s.Service.StoreSeller(r.Context(), seller)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (s *SellerHandler) UpdateSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["sellerId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		seller := &dto.SellerRequest{}
		if err := seller.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), seller); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := s.Service.UpdateSeller(r.Context(), parsedId, seller)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (s *SellerHandler) DeleteSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exists := mux.Vars(r)["sellerId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parseId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		if err := s.Service.DeleteSeller(r.Context(), parseId); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "seller deleted")
	}
}
