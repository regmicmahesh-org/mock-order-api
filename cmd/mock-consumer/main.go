package main

import (
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
	status := make(chan string)

	rabbitmq.Connect()
	rabbitmq.Initialize()
	wg.Add(1)
	log.Println("Connected.")

	go consumer.Receive(msgChannel)

	for {
		select {
		case msg := <-msgChannel:
			log.Println(msg.Contact, " <= OTP REQUEST")
			wg.Add(1)
			go twilio.SendOTP(msg, &wg, status)
		case sts := <-status:
			log.Println(sts)

		}
	}

}
