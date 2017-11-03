package main

import (
	"fmt"
	"log"
	"net/http"
)

var settings = Settings{
	currency: "USD",
	dbIP:     "192.168.1.1",
	dbPort:   "8072",
}

var ethWallet = "0xb837521FeE201B36bDAcdA38f6c0df97E0Cb6e9E"

func main() {
	log.Printf(settings.currency)
	result, err := getETHSellPriceKraken(settings.currency)
	if err == nil {
		log.Printf("result: %s", *result)
	}
}

func getETHSellPriceKraken(currency string) (*string, error) {

	// URL Request
	url := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=ETH%s", currency)

	// Build request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	var client *http.Client
	client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	result := "300"
	return &result, nil
	// var record GetMarkets

	// if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// for index := 0; index < len(record.Result); index++ {
	// 	// fmt.Println("Records :", record.Result[index])
	// 	// json, err := json.Marshal(record.Result[index])

	// 	resultJSON, err := json.Marshal(record.Result[index])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, err
	// 	}

	// 	fmt.Println(string(resultJSON))
	// }

	// return record.Result[0], nil
}
