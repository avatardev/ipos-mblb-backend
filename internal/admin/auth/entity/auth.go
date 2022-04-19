package entity

import (
	"database/sql"
)

type AuthLevelCtxKey string

type UserData struct {
	Id           int64
	Username     string
	Password     string
	Role         int64
	RoleName     string
	SellerID     *int64
	VehiclePlate *string
}

func (u *UserData) FromSql(row *sql.Row) error {
	return row.Scan(&u.Id, &u.Username, &u.Password, &u.Role, &u.RoleName, &u.SellerID, &u.VehiclePlate)
}
