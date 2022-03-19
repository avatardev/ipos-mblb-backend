package entity

import "time"

type Category struct {
	Id      uint64     `db:"id"`
	Name    string     `db:"nama_kategori"`
	Pajak   int64      `db:"pajak"`
	Status  bool       `db:"status"`
	Deleted *time.Time `db:"deleted_at"`
	Created *time.Time `db:"created_at"`
	Updated *time.Time `db:"updated_at"`
}

type Categories []*Category
