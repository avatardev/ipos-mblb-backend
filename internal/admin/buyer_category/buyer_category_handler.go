package buyercategory

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

type BuyerCategoryHandler struct {
	Service BuyerCategoryService
}

func NewBuyerCategoryHandler(service BuyerCategoryService) *BuyerCategoryHandler {
	return &BuyerCategoryHandler{Service: service}
}

func (b *BuyerCategoryHandler) GetCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1, 4) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetCategory] error: %v\n", err)
			if limit == "" {
				limitParsed = 0
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetCategory] error: %v\n", err)
			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")

		res, err := b.Service.GetCategory(r.Context(), keyword, limitParsed, offsetParsed)
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

func (b *BuyerCategoryHandler) GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1, 4) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		categoryId, exists := mux.Vars(r)["categoryId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrNotFound)
			return
		}

		parsedId, err := strconv.ParseInt(categoryId, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		res, err := b.Service.GetCategoryById(r.Context(), parsedId)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrNotFound)
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerCategoryHandler) StoreCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1, 4) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		category := &dto.BuyerCategoryRequest{}
		if err := category.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), category); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := b.Service.StoreCategory(r.Context(), category)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerCategoryHandler) UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1, 4) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		categoryId, exists := mux.Vars(r)["categoryId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(categoryId, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		category := &dto.BuyerCategoryRequest{}
		if err := category.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), category); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", err)
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := b.Service.UpdateCategory(r.Context(), parsedId, category)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerCategoryHandler) DeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1, 4) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		categoryId, exists := mux.Vars(r)["categoryId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(categoryId, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		if err := b.Service.DeleteCategory(r.Context(), parsedId); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "category deleted")
	}
}
