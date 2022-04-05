package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/dashboard/entity"

type DashboardInfoResponse struct {
	TrxCount    int64 `json:"trx_count"`
	SellerCount int64 `json:"seller_count"`
	BuyerCount  int64 `json:"buyer_count"`
	TaxSum      int64 `json:"tax_total"`
}

func NewDashboardInfoResponse(d *entity.DashboardInfo) *DashboardInfoResponse {
	return &DashboardInfoResponse{
		TrxCount:    d.TrxCount,
		SellerCount: d.SellerCount,
		BuyerCount:  d.BuyerCount,
		TaxSum:      d.TotalTax,
	}
}
