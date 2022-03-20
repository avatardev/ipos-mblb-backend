package category

import (
	"net/http"
	"strconv"

	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
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

func NewCategoryHandler(service CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}
