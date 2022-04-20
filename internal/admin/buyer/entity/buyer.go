package entity

type Buyer struct {
	VehiclePlate        string  `db:"plat_truk"`
	VehicleCategoryName string  `db:"nama_kategori"`
	VehicleCategoryId   uint64  `db:"kategori"`
	Company             string  `db:"perusahaan"`
	Phone               *string `db:"telp"`
	Address             string  `db:"alamat"`
	Email               string  `db:"email"`
	PICName             string  `db:"nama_pic"`
	PICPhone            *string `db:"hp_pic"`
	Description         string  `db:"keterangan"`
	Status              bool    `db:"status"`
}

type Buyers []*Buyer

type BuyerCompany struct {
	VPlate  string
	Company string
}

type BuyersCompany []*BuyerCompany
