package entity

import "time"

type Product struct {
	Id           int64   `db:"id"`
	CategoryName string  `db:"nama_kategori"`
	CategoryId   int64   `db:"id_kategori"`
	Name         string  `db:"nama_produk"`
	Price        float32 `db:"harga_std_m3"`
	Tax          uint64  `db:"pajak"`
	Description  string  `db:"keterangan"`
	Status       bool    `db:"status"`
	Img          *string
	Deleted      *time.Time `db:"deleted_at"`
	Created      *time.Time `db:"created_at"`
	Updated      *time.Time `db:"updated_at"`
}

type Products []*Product
