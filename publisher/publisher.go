package publisher

import (
	"github.com/regmicmahesh/mock-order-api/order"
	"github.com/streadway/amqp"
)

func send(order *order.Order) error {

	conn, err := amqp.Dial("a")

}
