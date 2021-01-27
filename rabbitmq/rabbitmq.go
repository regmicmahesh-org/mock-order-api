package rabbitmq

import (
	"github.com/streadway/amqp"
)

var Rq RabbitMQ

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
	Rq = RabbitMQ{Connection: conn, Channel: chann}
	return nil
}

func Initialize() error {

	_, err := Rq.Channel.QueueDeclare("sms", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil

}

func GetInstance() *RabbitMQ {
	return &Rq
}
