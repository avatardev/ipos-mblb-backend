package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/global/config"
)

type DatabaseClient struct {
	*sql.DB
}

func Init() *DatabaseClient {
	conf := config.GetConfig()

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Asia%%2FMakassar",
			conf.DBUsername,
			conf.DBPassword,
			conf.DBAddress,
			conf.DBName,
		))
	if err != nil {
		log.Fatalf("[Init] error on connecting to db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[Init] error pinging to db: %v", err)
	}

	return &DatabaseClient{DB: db}
}
