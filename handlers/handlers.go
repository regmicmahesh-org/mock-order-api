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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		return
	}

	ord, err := order.FromJSON(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("<=  OTP Request")
	log.Println("Publishing =>", ord.Contact)

	err = publisher.Send(ord)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Failed publishing =>", ord.Contact)
		return
	}

	log.Println("Succesfully published =>", ord.Contact)
	fmt.Fprintln(w, "OK")
}
