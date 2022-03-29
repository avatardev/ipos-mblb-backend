package order

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
)

type OrderHandler struct {
	Service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (o *OrderHandler) GenerateDetailTrx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), 1) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		query := r.URL.Query()

		startDate := query.Get("startDate")
		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Printf("[GenerateDetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		endDate := query.Get("endDate")
		endParsed, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			log.Printf("[GenerateDetailTrx] error: %v\n", err)
			responseutil.WriteErrorResponse(w, errors.ErrInvalidRequestBody)
			return
		}

		res, err := o.Service.GenerateDetailTrx(r.Context(), startParsed, endParsed)
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}
		filename := fmt.Sprintf("ExportDetailPeriode%v-%v", startDate, endDate)
		responseutil.WriteFileResponse(w, http.StatusCreated, filename, *res)
	}
}
