package entity

import (
	"database/sql"
)

type UserData struct {
	Id       int64
	Username string
	Password string
	Role     int64
}

func (u *UserData) FromSql(row *sql.Row) error {
	return row.Scan(&u.Id, &u.Username, &u.Password, &u.Role)
}
