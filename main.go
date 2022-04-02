package main

import (
	"database/sql"
	"goProject/api"
	db "goProject/db/sqlc"
	"goProject/until"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := until.LoadConfig(".")

	if err != nil {
		log.Fatal("connot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
