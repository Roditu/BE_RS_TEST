package main

import (
	"database/sql"
	"log"

	"github.com/Roditu/BE_RS_TEST/api"
	db "github.com/Roditu/BE_RS_TEST/db/sqlc"
	"github.com/Roditu/BE_RS_TEST/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDRiver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	tokenMaker := util.NewJWTMaker("supersecretkey")

	store := db.NewStore(conn)
	server := api.NewServer(store, tokenMaker)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}