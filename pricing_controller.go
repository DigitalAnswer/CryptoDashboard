package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := pc.client.Get(url)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var data map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &data); err == nil {
		result, _ := data["result"].(map[string]interface{})
		crypto := result["XETHZUSD"].(map[string]interface{})
		lastTrade := crypto["c"].([]interface{})
		lastPrice := lastTrade[0].(string)
		// fmt.Printf("%T\n", lastTrade)
		fmt.Println(lastPrice)
		return &lastPrice, nil
	}
	return nil, nil
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
