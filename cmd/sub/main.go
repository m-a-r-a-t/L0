package main

import (
	"fmt"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "2")
	c := make(chan string)

	go func() {
		exit := make(chan int)
		sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
			fmt.Printf("Received a message: %s\n", string(m.Data))
			c <- "Done"
			m.Ack()
		}, stan.DurableName("my-durable"), stan.SetManualAckMode(), stan.MaxInflight(100))

		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}
		<-exit
	}()

	for elem := range c {
		fmt.Println(elem)
	}

	// fmt.Println(<-c)

	sc.Close()
}
