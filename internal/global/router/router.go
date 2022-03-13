package router

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/buyer"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {

	buyerService := buyer.NewBuyerService()
	buyerHandler := buyer.NewBuyerHandler(buyerService)

	r.HandleFunc(AdminPing, buyerHandler.PingBuyer()).Methods(http.MethodGet)
	r.HandleFunc(AdminPingError, buyerHandler.PingError()).Methods(http.MethodGet)
}
