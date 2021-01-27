package publisher

import (
	"encoding/json"

	"github.com/regmicmahesh-org/mock-order-api/rabbitmq"

	"github.com/regmicmahesh-org/mock-order-api/order"
	"github.com/streadway/amqp"
)

func Send(order *order.Order) error {

	body, err := json.Marshal(*order)
	if err != nil {
		return err
	}

	err = rabbitmq.Rq.Channel.Publish("", "sms", false, false, amqp.Publishing{ContentType: "text/json", Body: []byte(body)})

	if err != nil {
		return err
	}

	return nil

}
