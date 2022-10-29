package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/m-a-r-a-t/L0/internal/http_server/models"
	stan "github.com/nats-io/stan.go"
)

func main() {

	jsonFile, err := os.Open("model.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	var order models.Order

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &order)

	sc, _ := stan.Connect("test-cluster", "1")

	err = sc.Publish("orders", byteValue)

	if err != nil {
		panic(err)
	}

	sc.Close()
}
