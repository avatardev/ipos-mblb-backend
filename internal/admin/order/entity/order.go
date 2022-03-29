package entity

import (
	"database/sql"
	"time"
)

// TODO Generate CSV REPORT (Date Range, No (auto), Order Date, Seller, Buyer, Payment, Qty, Status, Disc, Tax, Unit Price)

type TrxDetail struct {
	OrderDate       time.Time
	Company         string // Seller
	VehiclePlate    string // Buyer
	Payment         string
	ProductId       int64
	ProductIdUpdate int64
	Product         string
	Status          string
	Qty             int64
	QtyUpdate       int64
	Disc            int64
	Tax             int64
	TaxUpdate       int64
	Price           int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TrxDetails []*TrxDetail

func (t *TrxDetail) FromSql(row *sql.Rows) error {
	return row.Scan(&t.OrderDate, &t.Company, &t.VehiclePlate, &t.Payment, &t.ProductId, &t.ProductIdUpdate, &t.Qty, &t.QtyUpdate,
		&t.Status, &t.Disc, &t.Tax, &t.TaxUpdate, &t.Price, &t.CreatedAt, &t.UpdatedAt)
}
