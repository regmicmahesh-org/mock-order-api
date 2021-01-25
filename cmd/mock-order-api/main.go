package main

import (
	"net/http"

	"github.com/regmicmahesh/mock-order-api/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/order", handlers.OrderHandler)
	http.ListenAndServe(":8000", mux)
}
