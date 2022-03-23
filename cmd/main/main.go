package main

import (
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/global/config"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
	"github.com/avatardev/ipos-mblb-backend/internal/global/router"
	"github.com/avatardev/ipos-mblb-backend/pkg/middleware"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/httputil"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("[main] starting server")
	config.Init()

	conf := config.GetConfig()
	m := mux.NewRouter()
	db := database.Init()
	db.Ping()

	m.Use(middleware.CorsMiddleware())
	router.Init(m, db)

	httputil.ListenAndServe(conf.ServerAddress, m)
	log.Println("[main] stopping server gracefully")
}
