package main

import (
	"fmt"
	"net/http"

	"github.com/regmicmahesh-org/mock-order-api/handlers"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq"
)

func main() {
	err := rabbitmq.Connect()
	if err != nil {
		panic(err)
	}
	err = rabbitmq.Initialize()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/order", handlers.OrderHandler)
	fmt.Println("Starting Server at 8000")
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Server running at port 8000.")
	}

}
