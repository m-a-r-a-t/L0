package main

import (
	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "1")

	// Simple Synchronous Publisher
	err := sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	if err != nil {
		panic(err)
	}

	// Simple Async Subscriber
	// sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// })

	// // Unsubscribe
	// sub.Unsubscribe()

	// // Close connection
	sc.Close()
}
