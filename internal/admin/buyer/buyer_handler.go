package buyer

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
)

type BuyerHandler struct {
	Service BuyerService
}

func (b *BuyerHandler) PingBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := b.Service.Ping(r.Context())
		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}

func (b *BuyerHandler) PingError() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := b.Service.PingError(r.Context())
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
		}
	}
}

func NewBuyerHandler(service BuyerService) *BuyerHandler {
	return &BuyerHandler{Service: service}
}
