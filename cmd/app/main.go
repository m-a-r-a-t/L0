package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/m-a-r-a-t/L0/internal/http_server"
)

func main() {
	userName, password, dbName, host, port := "marat", "marat123", "marat", "localhost", 5432
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", userName, password, dbName, host, port)

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	ordersCache := map[string][]byte{}
	s := http_server.InitHttpServer(db, &ordersCache)
	s.Server.Listen(":3000")

}
