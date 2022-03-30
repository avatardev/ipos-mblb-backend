package entity

import "database/sql"

type Location struct {
	Id   int64
	Name string
}

type Locations []*Location

func (lc *Location) FromSql(row *sql.Row) error {
	return row.Scan(&lc.Id, &lc.Name)
}

func NewLocations(rows *sql.Rows) (Locations, error) {
	ls := Locations{}

	for rows.Next() {
		temp := &Location{}
		if err := rows.Scan(&temp.Id, &temp.Name); err != nil {
			return nil, err
		}
		ls = append(ls, temp)
	}
	return ls, nil
}
