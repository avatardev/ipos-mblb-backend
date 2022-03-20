package category

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

type CategoryHandler struct {
	Service CategoryService
}

func (c *CategoryHandler) GetCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := c.Service.GetCategory(r.Context())
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
		}

		if res == nil {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidResources)
			return
		}
		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (c *CategoryHandler) GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["categoryId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := c.Service.GetCategoryById(r.Context(), parsedId)
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

func (c *CategoryHandler) StoreCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := &dto.CategoryRequest{}
		if err := category.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), category); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		if err := c.Service.StoreCategory(r.Context(), category); err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, nil)
	}
}

func NewCategoryHandler(service CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}
