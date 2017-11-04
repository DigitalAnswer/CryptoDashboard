package main

import (
	"fmt"
	"log"
	"net/http"
)

// PricingController struct
type PricingController struct {
	client *http.Client
}

func (pc PricingController) getSellPriceKraken(cryptoCurrency string, fiatCurrency string) (*string, error) {
	// URL Request
	url := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=ETH%s", fiatCurrency)

	// Build request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := pc.client.Do(req)
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
