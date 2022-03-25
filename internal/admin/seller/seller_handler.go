package seller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
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
