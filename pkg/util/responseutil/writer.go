package responseutil

import (
	"log"
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/jsonutil"
)

func WriteSuccessResponse(rw http.ResponseWriter, status int, data interface{}) {
	BaseResponseWriter(rw, status, data, nil)
}

func WriteErrorResponse(rw http.ResponseWriter, err error) {
	errMetadata := errors.GetErrorResponseMetadata(err)
	BaseResponseWriter(rw, errMetadata.Status, nil, &dto.ErrorData{Code: errMetadata.Code, Message: errMetadata.Message})
}

func BaseResponseWriter(rw http.ResponseWriter, status int, data interface{}, er *dto.ErrorData) {
	res := dto.BaseResponse{Data: data, Error: er}
	jsonData, err := jsonutil.ConvertToJSON(res)
	if err != nil {
		log.Printf("[WriteSuccessResponse] json conversion error: %v", err)
		return
	}

	rw.WriteHeader(status)
	rw.Write(jsonData)
}
