package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/regmicmahesh-org/mock-order-api/order"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq/publisher"
)

//OrderHandler parses the json request received on body.
//Returns Ok if the response is good else returns
//the error.
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	ord, err := order.FromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%+v\n", ord)

	err = publisher.Send(ord)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "OK")
}
