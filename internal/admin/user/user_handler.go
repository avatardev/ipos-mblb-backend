package user

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

type UserHandler struct {
	Service UserService
}

const (
	USER_ADMIN   int64 = 1
	USER_CHECKER int64 = 4
)

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (u *UserHandler) GetUserAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetUserAdmin] limit: %v, error: %v\n", limit, err)

			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetUserAdmin] offset: %v, error: %v\n", offset, err)

			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")
		res, err := u.Service.GetUser(r.Context(), USER_ADMIN, keyword, limitParsed, offsetParsed)
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

func (u *UserHandler) GetUserChecker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := query.Get("limit")
		limitParsed, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			log.Printf("[GetUserAdmin] limit: %v, error: %v\n", limit, err)

			if limit == "" {
				limitParsed = 10
			}
		}

		offset := query.Get("offset")
		offsetParsed, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			log.Printf("[GetUserAdmin] offset: %v, error: %v\n", offset, err)

			if offset == "" {
				offsetParsed = 0
			}
		}

		keyword := query.Get("keyword")
		res, err := u.Service.GetUser(r.Context(), USER_CHECKER, keyword, limitParsed, offsetParsed)
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

func (u *UserHandler) GetUserAdminById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["userId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := u.Service.GetUserById(r.Context(), USER_ADMIN, parsedId)
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

func (u *UserHandler) GetUserCheckerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["userId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		res, err := u.Service.GetUserById(r.Context(), USER_CHECKER, parsedId)
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

func (u *UserHandler) StoreUserAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &dto.UserPostRequest{}
		if err := user.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), user); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := u.Service.StoreUser(r.Context(), USER_ADMIN, user)
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

func (u *UserHandler) StoreUserChecker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &dto.UserPostRequest{}
		if err := user.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), user); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := u.Service.StoreUser(r.Context(), USER_CHECKER, user)
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

func (u *UserHandler) UpdateUserAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["userId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		user := &dto.UserPutRequest{}
		if err := user.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), user); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := u.Service.UpdateUser(r.Context(), USER_ADMIN, parsedId, user)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (u *UserHandler) UpdateUserChecker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exist := mux.Vars(r)["userId"]
		if !exist {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		user := &dto.UserPutRequest{}
		if err := user.FromJSON(r.Body); err != nil {
			log.Printf("[FromJSON] error: %v\n", err)
			responseutil.WriteErrorResponse(w, err)
			return
		}

		v := validator.New()
		if err := v.StructCtx(r.Context(), user); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Printf("[Validation Error] error: %v\n", e.Field())
			}
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := u.Service.UpdateUser(r.Context(), USER_CHECKER, parsedId, user)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (u *UserHandler) DeleteUserAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exists := mux.Vars(r)["userId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		if err := u.Service.DeleteUser(r.Context(), USER_ADMIN, parsedId); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "user deleted")
	}
}

func (u *UserHandler) DeleteUserChecker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, exists := mux.Vars(r)["userId"]
		if !exists {
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responseutil.WriteErrorResponse(w, errors.ErrUnknown)
			return
		}

		if err := u.Service.DeleteUser(r.Context(), USER_CHECKER, parsedId); err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, "user deleted")
	}
}
