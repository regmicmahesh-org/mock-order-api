package main

import (
	"log"
	"net/http"
	"os"

	"github.com/regmicmahesh-org/mock-order-api/handlers"
	"github.com/regmicmahesh-org/mock-order-api/rabbitmq"
)

func main() {
	log.Println("Connecting to RabbitMQ")
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
	log.Println("Server => " + os.Args[1])
	log.Println("Endpoints => /order")
	err = http.ListenAndServe(os.Args[1], mux)
	if err != nil {
		panic(err)
	}
}
