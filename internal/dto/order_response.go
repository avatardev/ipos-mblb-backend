package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/order/entity"
)

type TrxDetail struct {
	OrderDate    string  `json:"order_date"`
	Company      string  `json:"seller"`
	VehiclePlate string  `json:"buyer"`
	Payment      string  `json:"payment_method"`
	Status       string  `json:"status"`
	Product      string  `json:"product_name"`
	Note         *string `json:"note"`
	OrderId      int64   `json:"order_id"`
	Qty          int64   `json:"qty"`
	Disc         int64   `json:"disc"`
	Tax          int64   `json:"tax"`
	Price        int64   `json:"price"`
}

type TrxDetails []*TrxDetail
type TrxDetailsJSON struct {
	TrxDetail []*TrxDetail `json:"trx_detail"`
	Limit     uint64       `json:"limit"`
	Offset    uint64       `json:"offset"`
	Total     uint64       `json:"total"`
}

type TrxBrief struct {
	OrderDate  string `json:"order_date"`
	Company    string `json:"company"`
	Buyer      string `json:"buyer"`
	OrderId    int64  `json:"order_id"`
	TotalTax   int64  `json:"total_tax"`
	TotalPrice int64  `json:"total_price"`
}

type TrxBriefs []*TrxBrief

type TrxBriefsJSON struct {
	TrxBrief []*TrxBrief `json:"trx_brief"`
	Limit    uint64      `json:"limit"`
	Offset   uint64      `json:"offset"`
	Total    uint64      `json:"total"`
}

func NewTrxDetail(trx *entity.TrxDetail) *TrxDetail {
	qty := trx.Qty
	if trx.QtyUpdate != 0 {
		qty = trx.QtyUpdate
	}

	tax := trx.Tax
	if trx.TaxUpdate != 0 {
		tax = trx.TaxUpdate
	}

	return &TrxDetail{
		OrderId:      trx.Orderid,
		OrderDate:    trx.OrderDate.Format("02/01/2006"),
		Company:      trx.Company,
		VehiclePlate: trx.VehiclePlate,
		Payment:      trx.Payment,
		Status:       trx.Status,
		Product:      trx.Product,
		Qty:          qty,
		Disc:         trx.Disc,
		Tax:          tax,
		Price:        trx.Price,
		Note:         trx.Note,
	}
}

func NewTrxDetailsJSON(data entity.TrxDetails, limit uint64, offset uint64, count uint64) *TrxDetailsJSON {
	res := &TrxDetailsJSON{
		Limit:  limit,
		Offset: offset,
		Total:  count,
	}

	res.TrxDetail = NewTrxDetails(data)
	return res
}

func NewTrxDetails(data entity.TrxDetails) TrxDetails {
	res := TrxDetails{}

	for _, trx := range data {
		trx := NewTrxDetail(trx)
		res = append(res, trx)
	}

	return res
}

func (t *TrxDetail) ToSlice(idx int) []string {
	return []string{
		strconv.Itoa(idx + 1),
		t.OrderDate,
		t.Company,
		t.VehiclePlate,
		t.Payment,
		t.Status,
		t.Product,
		strconv.FormatInt(t.Qty, 10),
		strconv.FormatInt(t.Disc, 10),
		strconv.FormatInt(t.Tax, 10),
		strconv.FormatInt(t.Price, 10),
	}
}

func (t *TrxDetails) ToCSV(dateStart time.Time, dateEnd time.Time) [][]string {
	res := [][]string{
		{"Rekapitulasi Penjualan MBLB", "", "", "", "", "", "", "", "", ""},
		{fmt.Sprintf("Periode: %v - %v", dateStart.Format("02/01/2006"), dateEnd.Format("02/01/2006")), "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"No", "Order Date", "Seller", "Buyer", "Payment", "Status", "Produk", "Qty", "Disc", "Tax", "Unit Price"},
	}

	for idx, data := range *t {
		res = append(res, data.ToSlice(idx))
	}
	return res
}

func NewTrxBriefsJSON(data entity.TrxDetails, limit uint64, offset uint64, count uint64) *TrxBriefsJSON {
	res := &TrxBriefsJSON{
		Limit:  limit,
		Offset: offset,
		Total:  count,
	}

	res.TrxBrief = NewTrxBriefs(data)
	return res
}

func NewTrxBriefs(data entity.TrxDetails) TrxBriefs {
	res := TrxBriefs{}

	for _, trx := range data {
		tax := trx.Tax
		if trx.TaxUpdate != 0 {
			tax = trx.TaxUpdate
		}

		res = append(res, &TrxBrief{
			OrderDate:  trx.OrderDate.Format("02/01/2006"),
			Company:    trx.Company,
			Buyer:      trx.VehiclePlate,
			OrderId:    trx.Orderid,
			TotalTax:   tax,
			TotalPrice: trx.Price,
		})
	}

	return res
}

func (t *TrxBrief) ToSlice(idx int) []string {
	return []string{
		strconv.Itoa(idx + 1),
		t.OrderDate,
		t.Company,
		t.Buyer,
		strconv.FormatInt(t.TotalTax, 10),
		strconv.FormatInt(t.TotalPrice, 10),
	}
}

func (td *TrxBriefs) ToCSV(dateStart time.Time, dateEnd time.Time) [][]string {
	res := [][]string{
		{"Rekapitulasi Penjualan MBLB", "", "", "", "", ""},
		{fmt.Sprintf("Periode: %v - %v", dateStart.Format("02/01/2006"), dateEnd.Format("02/01/2006")), "", "", "", "", ""},
		{"", "", "", "", "", ""},
		{"", "", "", "", "", ""},
		{"No", "Order Date", "Seller", "Buyer", "Total Tax", "Total Price"},
	}

	for idx, data := range *td {
		res = append(res, data.ToSlice(idx))
	}
	return res
}
