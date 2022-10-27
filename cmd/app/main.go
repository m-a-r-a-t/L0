package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/m-a-r-a-t/L0/internal/http_server"
	"github.com/m-a-r-a-t/L0/internal/http_server/models"

	// "github.com/m-a-r-a-t/L0/internal/http_server/repositories"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	userName, password, dbName, host, port := "marat", "marat123", "marat", "localhost", 5432
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", userName, password, dbName, host, port)

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ordersCache := map[string][]byte{}
	// orderRepo := repositories.NewOrderRepo(db)

	jsonFile, err := os.Open("model.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	var order models.Order

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &order)

	// orderRepo.InsertOrders([]*models.Order{&order})
	fmt.Println("_____________________________")
	s := http_server.InitHttpServer(db, &ordersCache)
	fmt.Println(s)

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect("1", "2", stan.NatsConn(nc))
	fmt.Println(err, sc)
	// Simple Synchronous Publisher
	// sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// s.Server.Listen(":3000")

}

func InitCache() {

}
