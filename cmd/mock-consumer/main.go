package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/regmicmahesh-org/mock-order-api/rabbitmq"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq/consumer"
	"github.com/regmicmahesh-org/mock-order-api/twilio"

	"github.com/regmicmahesh-org/mock-order-api/order"
)

var wg sync.WaitGroup

func main() {

	msgChannel := make(chan *order.Order)

	rabbitmq.Connect()
	wg.Add(1)

	go consumer.Receive(msgChannel)

	for {
		msg := <-msgChannel
		if err := twilio.SendOTP(msg); err != nil {
			fmt.Println(err)
			log.Printf("Failed to send message.")
		} else {
			fmt.Println("Succesfully sent message to =>", msg.Contact)
		}
	}
	wg.Wait()

}
