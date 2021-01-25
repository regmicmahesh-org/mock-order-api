package order

import (
	"encoding/json"
	"errors"
	"io"
)

type Order struct {
	Item    string
	Contact string
	Price   int
	Time    string
}

func FromJSON(jsonString io.Reader) (*Order, error) {
	var order Order
	err := json.NewDecoder(jsonString).Decode(&order)
	if err != nil {
		return nil, err
	}

	err = order.Validate()
	if err != nil {
		return nil, err
	}
	return &order, nil

}

//Validate is used to validate an order.
func (order *Order) Validate() error {
	if order.Item == "" {
		return errors.New("Item cannot be blank")
	}
	if len(order.Contact) != 10 {
		return errors.New("Invalid Phone Number")
	}
	if order.Price < 0 {
		return errors.New("Invalid Pricing")
	}
	return nil
}
