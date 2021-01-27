package consumer

import (
	"bytes"
	"sync"

	"github.com/regmicmahesh-org/mock-order-api/order"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq"
)

var wg sync.WaitGroup

func Receive(returnVal chan *order.Order) {

	rabbitmq.Connect()
	instance := rabbitmq.GetInstance()

	msgs, err := instance.Channel.Consume("sms", "", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	wg.Add(1)

	for {
		message := <-msgs
		var ord, err = order.FromJSON(bytes.NewReader(message.Body))
		if err != nil {
			panic(err)
		}
		returnVal <- ord

	}

}
