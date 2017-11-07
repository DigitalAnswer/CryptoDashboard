package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PricingController struct
type PricingController struct {
	client *http.Client
}

// CurrencyTickerResponse struct
type CurrencyTickerResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rank     string `json:"rank"`
	PriceUSD string `json:"price_usd"`
	PriceBTC string `json:"price_btc"`
	// "24h_volume_usd": "72855700.0",
	// "market_cap_usd": "9080883500.0",
	// "available_supply": "15844176.0",
	// "total_supply": "15844176.0",
	// "max_supply": "21000000.0",
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// cryptoCurrency must the full name of the crypto like bitcoin and not btc
func (pc PricingController) getSellPriceKraken(cryptoCurrency string, fiatCurrency string) (*CurrencyTickerResponse, error) {
	// URL Request
	// url := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=ETH%s", fiatCurrency)
	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s/", cryptoCurrency)

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

	var record []*CurrencyTickerResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
		return nil, err
	}

	return record[0], nil
}
