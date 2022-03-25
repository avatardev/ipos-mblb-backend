package entity

import "time"

type Seller struct {
	Id          int64
	Company     string
	Phone       string
	Address     string
	District    string
	Email       string
	PICName     string
	PICPhone    string
	NPWP        string
	KTP         string
	NoIUP       string
	ValidPeriod time.Time
	Description string
	Status      bool
}

type Sellers []*Seller
