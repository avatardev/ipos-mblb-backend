package errors

import (
	"errors"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/dto"
)

var (
	ErrBuyerPing          = errors.New("5f3467a8-a2b2-11ec-b909-0242ac120002")
	ErrUnknown            = errors.New("755cb64b-9d59-4df7-ad53-9275b9a2e6f6")
	ErrInvalidRequestBody = errors.New("84d48c05-b61f-42b5-98d9-e8d54a925df5")
	ErrInvalidResources   = errors.New("b73e5586-b5fa-4c30-92d4-da7d4c9675d8")
	ErrNotFound           = errors.New("8e8a54ae-c2f0-451d-80c9-24f71764f9e5")
	ErrUserPriv           = errors.New("d9614ed1-8534-443c-ae7e-192dfe933f75")
	ErrUserCredential     = errors.New("b5ebf348-2b34-4032-9363-74d7a4759466")
)

var errorMap = map[error]dto.ErrorResponseMetadata{
	ErrBuyerPing:          NewErrorResponseMetadata(http.StatusBadRequest, ErrBuyerPing.Error(), "just some dummy error bro"),
	ErrUnknown:            NewErrorResponseMetadata(http.StatusInternalServerError, ErrUnknown.Error(), "internal server error"),
	ErrInvalidRequestBody: NewErrorResponseMetadata(http.StatusBadRequest, ErrInvalidRequestBody.Error(), "invalid request body"),
	ErrInvalidResources:   NewErrorResponseMetadata(http.StatusNotFound, ErrInvalidResources.Error(), "resources is empty"),
	ErrNotFound:           NewErrorResponseMetadata(http.StatusNotFound, ErrNotFound.Error(), "resources not found"),
	ErrUserPriv:           NewErrorResponseMetadata(http.StatusForbidden, ErrUserPriv.Error(), "user doesn't have enough privilege to access this resources"),
	ErrUserCredential:     NewErrorResponseMetadata(http.StatusBadRequest, ErrUserCredential.Error(), "invalid user credentials"),
}

func NewErrorResponseMetadata(status int, code string, message string) dto.ErrorResponseMetadata {
	return dto.ErrorResponseMetadata{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func GetErrorResponseMetadata(err error) (er dto.ErrorResponseMetadata) {
	er, ok := errorMap[err]
	if !ok {
		er = errorMap[ErrUnknown]
	}
	return
}
