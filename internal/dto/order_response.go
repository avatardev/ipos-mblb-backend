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
}

type TrxDaily struct {
	OrderDate string          `json:"order_date"`
	Details   map[int64]int64 `json:"details"`
}

type TrxDailes []*TrxDaily

type TrxDailiesJSON struct {
	TrxDaily []*TrxDaily `json:"trx_daily"`
}

type TrxMonitor struct {
	OrderId      int64   `json:"order_id"`
	OrderDate    string  `json:"order_date"`
	VehiclePlate string  `json:"buyer"`
	Company      string  `json:"seller"`
	Product      string  `json:"product_name"`
	Qty          int64   `json:"qty"`
	VolumePrice  float64 `json:"volume_price"`
	SellPrice    float64 `json:"sell_price"`
	Tax          int64   `json:"tax"`
	TaxUpdate    int64   `json:"bpkad_tax"`
}

type TrxMonitors []*TrxMonitor

type TrxMonitorJSON struct {
	TrxMonitor []*TrxMonitor `json:"trx_monitor"`
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

func NewTrxDetailsJSON(data entity.TrxDetails) *TrxDetailsJSON {
	res := &TrxDetailsJSON{}

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

func NewTrxBriefsJSON(data entity.TrxDetails) *TrxBriefsJSON {
	res := &TrxBriefsJSON{}

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

func NewTrxDailiesJSON(data entity.TrxDailies) *TrxDailiesJSON {
	res := &TrxDailiesJSON{}

	res.TrxDaily = NewTrxDailies(data)
	return res
}

func NewTrxDailies(data entity.TrxDailies) TrxDailes {
	res := TrxDailes{}
	year, month, _ := data[0].Date.Date()
	daysInMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1).Day()

	for i := 1; i <= daysInMonth; i++ {
		res = append(res, &TrxDaily{
			OrderDate: time.Date(year, month, i, 0, 0, 0, 0, time.UTC).Format("02"),
			Details:   map[int64]int64{},
		})
	}

	// loop over trx
	for _, trx := range data {
		date := trx.Date.Format("02")

		for _, val := range res {
			if val.OrderDate == date {
				val.Details[trx.Volume] = trx.Quantity
				break
			}
		}
	}

	return res
}

func (t *TrxDaily) ToSlice(idx int) []string {
	var t1, t2, t3, t4, t5 string

	if _, ok := t.Details[6]; ok {
		t1 = strconv.FormatInt(t.Details[6], 10)
	}
	if _, ok := t.Details[7]; ok {
		t2 = strconv.FormatInt(t.Details[7], 10)
	}
	if _, ok := t.Details[8]; ok {
		t3 = strconv.FormatInt(t.Details[8], 10)
	}
	if _, ok := t.Details[10]; ok {
		t4 = strconv.FormatInt(t.Details[10], 10)
	}
	if _, ok := t.Details[20]; ok {
		t5 = strconv.FormatInt(t.Details[20], 10)
	}

	return []string{
		t.OrderDate,
		t1, t2, t3, t4, t5,
	}
}

func (t *TrxDailes) ToCSV(company string, npwp int64, month string, year int) [][]string {
	res := [][]string{
		{"LAPORAN HARIAN", ""},
		{fmt.Sprintf("Nama WP: %v", company), ""},
		{fmt.Sprintf("NPWPD: %v", npwp), ""},
		{fmt.Sprintf("MASA: %v %v", month, year), ""},
		{"", ""},
		{"", ""},
		{"Tanggal", "6", "7", "8", "10", "20"},
	}

	for idx, data := range *t {
		res = append(res, data.ToSlice(idx))
	}
	return res
}

func NewTrxMonitorJSON(data entity.TrxMonitors) *TrxMonitorJSON {
	res := &TrxMonitorJSON{}

	res.TrxMonitor = NewTrxMonitors(data)
	return res
}

func NewTrxMonitors(data entity.TrxMonitors) TrxMonitors {

	res := TrxMonitors{}

	for _, trx := range data {
		product := trx.MBLBType
		if trx.MBLBTypeUpdate != nil {
			product = *trx.MBLBTypeUpdate
		}

		qty := trx.Volume
		if trx.VolumeUpdate != 0 {
			qty = trx.VolumeUpdate
		}

		res = append(res, &TrxMonitor{
			OrderId:      trx.OrderId,
			OrderDate:    trx.OrderDate.Format("02/01/2006"),
			VehiclePlate: trx.VehiclePlate,
			Company:      trx.Company,
			Product:      product,
			Qty:          qty,
			VolumePrice:  trx.VolumePrice,
			SellPrice:    trx.SoldPrice,
			Tax:          trx.Tax,
			TaxUpdate:    trx.TaxUpdate,
		})
	}

	return res
}

func (t *TrxMonitor) ToSlice(idx int) []string {
	return []string{
		strconv.Itoa(idx + 1),
		t.OrderDate,
		t.Company,
		t.VehiclePlate,
		t.Product,
		strconv.FormatInt(t.Qty, 10),
		fmt.Sprintf("%.2f", t.VolumePrice),
		fmt.Sprintf("%.2f", t.SellPrice),
		strconv.FormatInt(t.Tax, 10),
		strconv.FormatInt(t.TaxUpdate, 10),
	}
}

func (t *TrxMonitors) ToCSV(dateStart time.Time, dateEnd time.Time) [][]string {
	res := [][]string{
		{"Rekapitulasi Penjualan MBLB", "", "", "", "", "", "", "", "", ""},
		{fmt.Sprintf("Periode: %v - %v", dateStart.Format("02/01/2006"), dateEnd.Format("02/01/2006")), "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"No", "Order Date", "Seller", "Buyer", "MBLB Material", "M3 Volume", "M3 Price", "Sell Price", "Tax", "Monitoring BPKAD"},
	}

	for idx, data := range *t {
		res = append(res, data.ToSlice(idx))
	}
	return res
}
