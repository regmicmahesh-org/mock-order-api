package main

import (
	"net/http"

	"github.com/regmicmahesh-org/mock-order-api/handlers"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq/publisher"
)

func main() {
	err := publisher.Connect()
	if err != nil {
		panic(err)
	}
	err = publisher.Initialize()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/order", handlers.OrderHandler)
	http.ListenAndServe(":8000", mux)
}
