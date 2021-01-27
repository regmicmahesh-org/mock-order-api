package publisher

import (
	"encoding/json"

	"github.com/regmicmahesh-org/mock-order-api/order"
	"github.com/streadway/amqp"
)

var rq RabbitMQ

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
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

	rq = RabbitMQ{Connection: conn, Channel: chann}
	return nil
}

func Initialize() error {

	_, err := rq.Channel.QueueDeclare("sms", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil

}

func GetInstance() *RabbitMQ {
	return &rq
}

func Send(order *order.Order) error {

	body, err := json.Marshal(*order)
	if err != nil {
		return err
	}

	err = rq.Channel.Publish("", "sms", false, false, amqp.Publishing{ContentType: "text/json", Body: []byte(body)})

	if err != nil {
		return err
	}

	return nil

}
