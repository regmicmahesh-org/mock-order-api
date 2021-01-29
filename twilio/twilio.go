package twilio

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/regmicmahesh-org/mock-order-api/order"
)

func prepareRequest(number string, body string) *http.Request {
	data := url.Values{}
	data.Set("To", "+977"+number)
	data.Set("MessagingServiceSid", "MG3eed970714f552e0ef7e0bac6b59890c")
	data.Set("Body", body)
	numberByte := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST",
		"https://api.twilio.com/2010-04-01/Accounts/AC2ee6f316d67fb729c7e2b53769770f36/Messages.json",
		numberByte)
	req.SetBasicAuth("AC2ee6f316d67fb729c7e2b53769770f36", "1be295374ae96a8ba6470947d5a017db")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func prepareBody(ord *order.Order) string {

	return "Your order for " + ord.Item + " was succesful. You need to pay " + strconv.Itoa(ord.Price) + ". Thank you!"

}

func SendOTP(ord *order.Order, wg *sync.WaitGroup, status chan string) {

	client := &http.Client{}
	body := prepareBody(ord)
	req := prepareRequest(ord.Contact, body)
	resp, err := client.Do(req)
	defer wg.Done()
	if resp.StatusCode != 201 || err != nil {
		status <- "OTP FAILED => " + ord.Contact
		return
	}

	status <- "OTP SENT => " + ord.Contact

}
