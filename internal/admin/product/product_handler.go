package product

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
)

type ProductHandler struct {
	Service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (p *ProductHandler) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := p.Service.GetProduct(r.Context())
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
