package errors

import (
	"errors"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/dto"
)

var (
	ErrBuyerPing = errors.New("5f3467a8-a2b2-11ec-b909-0242ac120002")
	ErrUnknown   = errors.New("755cb64b-9d59-4df7-ad53-9275b9a2e6f6")
)

var errorMap = map[error]dto.ErrorResponseMetadata{
	ErrBuyerPing: NewErrorResponseMetadata(http.StatusBadRequest, ErrBuyerPing.Error(), "just some dummy error bro"),
	ErrUnknown:   NewErrorResponseMetadata(http.StatusInternalServerError, ErrUnknown.Error(), "internal server error"),
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
