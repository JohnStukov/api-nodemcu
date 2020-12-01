package routers

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Txt(dispositivo string, tipo string) {
	// Set initial variables
	accountSid := "ACfe54870522cae6697bd9e6337bfa7cbb"
	authToken := "edeae3f09b4a806e28765cb44267f8f4"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To","+524271417372")
	v.Set("From","+12056977554")
	v.Set("Body","Se disparo la alarma " + tipo + "! Dispositivo: " + dispositivo)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	log.Println(resp.Status, "mensaje enviado!")
}