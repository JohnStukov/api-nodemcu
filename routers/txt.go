package routers

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Txt(dispositivo string, tipo string) {
	// Set initial variables
	accountSid := ""
	authToken := ""
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To","")
	v.Set("From","")
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
