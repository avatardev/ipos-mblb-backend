package entity

import (
	"database/sql"
	"time"
)

type LogCtxKey string
type LogInfo struct {
	Id        int64
	UserId    int64
	Username  *string
	Message   string
	Timestamp time.Time
}

type Logs []*LogInfo

func NewLogs(rows *sql.Rows) (Logs, error) {
	lg := Logs{}

	for rows.Next() {
		temp := &LogInfo{}
		if err := rows.Scan(&temp.Id, &temp.UserId, &temp.Username, &temp.Message, &temp.Timestamp); err != nil {
			return nil, err
		}
		lg = append(lg, temp)
	}

	return lg, nil
}
