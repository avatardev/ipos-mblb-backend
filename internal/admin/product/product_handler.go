package product

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

func (p *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, exists := mux.Vars(r)["productId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(productId, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := p.Service.GetProductById(r.Context(), parsedId)
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

func (p *ProductHandler) StoreProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product := &dto.ProductRequest{}
		if err := product.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), product); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := p.Service.StoreProduct(r.Context(), product)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
