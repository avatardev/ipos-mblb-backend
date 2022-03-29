package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/order/entity"
)

type TrxDetail struct {
	OrderDate    time.Time
	Company      string // Seller
	VehiclePlate string // Buyer
	Payment      string
	Status       string
	Product      string
	Qty          int64
	Disc         int64
	Tax          int64
	Price        int64
}

type TrxDetails []*TrxDetail

func NewTrxDetails(data entity.TrxDetails) TrxDetails {
	res := TrxDetails{}

	for _, trx := range data {
		qty := trx.Qty
		if trx.QtyUpdate != 0 {
			qty = trx.QtyUpdate
		}

		tax := trx.Tax
		if trx.TaxUpdate != 0 {
			tax = trx.TaxUpdate
		}

		res = append(res, &TrxDetail{
			OrderDate:    trx.OrderDate,
			Company:      trx.Company,
			VehiclePlate: trx.VehiclePlate,
			Payment:      trx.Payment,
			Status:       trx.Status,
			Product:      trx.Product,
			Qty:          qty,
			Disc:         trx.Disc,
			Tax:          tax,
			Price:        trx.Price,
		})
	}

	return res
}

func (t *TrxDetail) ToSlice(idx int) []string {
	return []string{
		strconv.Itoa(idx + 1),
		t.OrderDate.Format("02/01/2006"),
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

func (td *TrxDetails) ToCSV(dateStart time.Time, dateEnd time.Time) [][]string {
	res := [][]string{
		{"Rekapitulasi Penjualan MBLB", "", "", "", "", "", "", "", "", ""},
		{fmt.Sprintf("Periode: %v - %v", dateStart.Format("02/01/2006"), dateEnd.Format("02/01/2006")), "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"No", "Order Date", "Seller", "Buyer", "Payment", "Status", "Produk", "Qty", "Disc", "Tax", "Unit Price"},
	}

	for idx, data := range *td {
		res = append(res, data.ToSlice(idx))
	}
	return res
}
