package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type payload struct {
	Endpoint      string                   `json:"endpoint"`
	Quotes        []map[string]interface{} `json:"quotes"`
	RequestedTime string                   `json:"requested_time"`
	Timestamp     int                      `json:"timestamp"`
}

func main() {

	currencies := "EURUSD,GBPUSD"
	apiKey := "your api"

	URL := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + apiKey

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println()
		if err != nil {
			log.Fatal(err)
		}
		data := payload{}
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		fmt.Println("Endpoint:", data.Endpoint, "\nRequested Time:", data.RequestedTime, "\nTimestamp:", data.Timestamp)

		for _, v := range data.Quotes {
			fmt.Println("ask:", v["ask"])
			fmt.Println("base_currency:", v["base_currency"])
			fmt.Println("bid:", v["bid"])
			fmt.Println("mid:", v["mid"])
			fmt.Println("quote_currency:", v["quote_currency"])
		}
	}
}
