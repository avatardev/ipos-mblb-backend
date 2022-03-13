package main

import (
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/global/config"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/httputil"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("[main] starting server")
	config.Init()

	conf := config.GetConfig()
	router := mux.NewRouter()
	db := database.Init()
	db.Ping()

	httputil.ListenAndServe(conf.ServerAddress, router)
	log.Println("[main] stopping server gracefully")
}
