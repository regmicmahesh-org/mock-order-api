package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/regmicmahesh-org/mock-order-api/twilio"

	"github.com/regmicmahesh-org/mock-order-api/order"

	"github.com/regmicmahesh-org/mock-order-api/consumer"
	"github.com/regmicmahesh-org/mock-order-api/publisher"
)

var wg sync.WaitGroup

func main() {

	msgChannel := make(chan *order.Order)

	publisher.Connect()
	wg.Add(1)

	go consumer.Receive(msgChannel)

	for {
		msg := <-msgChannel
		if err := twilio.SendOTP(msg); err != nil {
			fmt.Println(err)
			log.Printf("Failed to send message.")
		}
	}
	wg.Wait()

}