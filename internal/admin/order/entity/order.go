package entity

import (
	"database/sql"
	"time"
)

type TrxDetail struct {
	Company         string // Seller
	VehiclePlate    string // Buyer
	Payment         string
	Status          string
	Product         string
	Note            *string
	Orderid         int64
	ProductId       int64
	ProductIdUpdate int64
	Qty             int64
	QtyUpdate       int64
	Disc            int64
	Tax             int64
	TaxUpdate       int64
	Price           int64
	OrderDate       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TrxDetails []*TrxDetail

func (t *TrxDetail) FromSql(row *sql.Rows) error {
	return row.Scan(&t.Orderid, &t.OrderDate, &t.Company, &t.VehiclePlate, &t.Payment, &t.ProductId, &t.ProductIdUpdate, &t.Qty, &t.QtyUpdate,
		&t.Status, &t.Disc, &t.Tax, &t.TaxUpdate, &t.Price, &t.Note, &t.CreatedAt, &t.UpdatedAt)
}

func (t *TrxDetail) FromSingleSql(row *sql.Row) error {
	return row.Scan(&t.Orderid, &t.OrderDate, &t.Company, &t.VehiclePlate, &t.Payment, &t.ProductId, &t.ProductIdUpdate, &t.Qty, &t.QtyUpdate,
		&t.Status, &t.Disc, &t.Tax, &t.TaxUpdate, &t.Price, &t.Note, &t.CreatedAt, &t.UpdatedAt)
}

type TrxMonitor struct {
	MBLBTypeUpdate *string
	Company        string
	VehiclePlate   string
	MBLBType       string
	OrderId        int64
	Volume         int64
	VolumePrice    float64
	SoldPrice      float64
	Tax            int64
	VolumeUpdate   int64
	TaxUpdate      int64
	OrderDate      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TrxMonitors []*TrxMonitor

func (m *TrxMonitor) FromOrderSql(row *sql.Rows) error {
	return row.Scan(&m.OrderId, &m.OrderDate, &m.Company, &m.VehiclePlate, &m.MBLBType, &m.MBLBTypeUpdate, &m.Volume, &m.VolumePrice, &m.SoldPrice, &m.VolumeUpdate, &m.Tax, &m.TaxUpdate, &m.CreatedAt, &m.UpdatedAt)
}

type TrxDaily struct {
	Date     time.Time
	Volume   int64
	Quantity int64
}

type TrxDailies []*TrxDaily

func (m *TrxDaily) FromSql(row *sql.Rows) error {
	return row.Scan(&m.Date, &m.Volume, &m.Quantity)
}
