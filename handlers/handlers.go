package handlers

import (
	"fmt"
	"net/http"

	"github.com/regmicmahesh/mock-order-api/order"
)

//OrderHandler parses the json request received on body.
//Returns Ok if the response is good else returns
//the error.
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	_, err := order.FromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "OK")
}
