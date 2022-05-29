package product

import (
	"bytes"
	"io"
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

type ProductHandler struct {
	Service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (p *ProductHandler) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetProduct] error: %v\n", err)
			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offestParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetProduct] error: %v\n", err)
			if offset == "" {
				offestParsed = 0
			}
		}

		keyword := query.Get("keyword")

		res, err := p.Service.GetProduct(r.Context(), keyword, limitParsed, offestParsed)
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
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

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
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

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
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (p *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

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

		res, err := p.Service.UpdateProduct(r.Context(), parsedId, product)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (p *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		productId, exists := mux.Vars(r)["productId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(productId, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		if err := p.Service.DeleteProduct(r.Context(), parsedId); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "product deleted")
	}
}

func (p *ProductHandler) EditProductImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		id := mux.Vars(r)["productId"]
		parsedID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Printf("[EditProductImage] invalid product ID, id => %s, err => %+v\n", id, err)
			panic(errors.ErrInvalidRequestBody)
		}

		imgData, meta, err := r.FormFile("img-data")
		if err != nil {
			log.Panicf("[EditProductImage] failed to read img data, err => %+v\n", err)
		}

		defer imgData.Close()

		buff := new(bytes.Buffer)
		_, err = io.Copy(buff, imgData)
		if err != nil {
			log.Panicf("[EditProductImage] failed to copy img data, err => %+v\n", err)
		}

		res, err := p.Service.EditProductImage(r.Context(), parsedID, buff, meta.Filename)
		if err != nil {
			log.Panicf("[EditProductImage] an error occured while processing product's img, id => %d, %+v\n", parsedID, err)
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
