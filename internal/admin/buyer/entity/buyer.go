package entity

type Buyer struct {
	VehiclePlate        string `db:"plat_truk"`
	VehicleCategoryName string `db:"nama_kategori"`
	VehicleCategoryId   uint64 `db:"kategori"`
	Company             string `db:"perusahaan"`
	Phone               string `db:"telp"`
	Address             string `db:"alamat"`
	Email               string `db:"email"`
	PICName             string `db:"nama_pic"`
	PICPhone            string `db:"hp_pic"`
	Description         string `db:"keterangan"`
	Status              bool   `db:"status"`
}

type Buyers []*Buyer

type Category struct {
	Id           uint64 `db:"id"`
	Name         string `db:"nama_kategori"`
	MultiProduct bool   `db:"multi_produk"`
}

type Categories []*Category
