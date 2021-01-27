package twilio

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	req.SetBasicAuth("AC2ee6f316d67fb729c7e2b53769770f36", "b698f3bc0d81f0da6ff3ff5760b50655")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func prepareBody(ord *order.Order) string {

	return "Your order for " + ord.Item + " was succesful. You need to pay " + strconv.Itoa(ord.Price) + ". Thank you!"

}

func SendOTP(ord *order.Order) error {

	client := &http.Client{}
	body := prepareBody(ord)
	req := prepareRequest(ord.Contact, body)
	resp, err := client.Do(req)
	if resp.StatusCode != 201 {
		return errors.New(resp.Status)
	}
	if err != nil {
		return err
	}

	return nil

}
