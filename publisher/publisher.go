package publisher

import (
	"encoding/json"
	"fmt"

	"github.com/regmicmahesh-org/mock-order-api/order"
	"github.com/streadway/amqp"
)

var rq RabbitMQ

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func init() {
	fmt.Println("hi")
}

func Connect() error {

	conn, err := amqp.Dial("amqp://guest:guest@10.6.0.7:5672")
	if err != nil {
		return err
	}
	chann, err := conn.Channel()
	if err != nil {
		return err
	}

	rq = RabbitMQ{connection: conn, channel: chann}
	return nil
}

func Initialize() error {

	_, err := rq.channel.QueueDeclare("sms", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil

}

func Send(order *order.Order) error {

	body, err := json.Marshal(*order)
	if err != nil {
		return err
	}

	err = rq.channel.Publish("", "sms", false, false, amqp.Publishing{ContentType: "text/json", Body: []byte(body)})

	if err != nil {
		return err
	}

	return nil

}
