package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// APIPath enum string
type APIPath string

// APIPath enum string
const (
	APIPathTicker APIPath = "ticker"
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

func (r *CurrencyTickerResponse) UnmarshalJSON(raw []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}

	if v, ok := data["id"]; ok {
		r.ID = getStringValue(v)
	}
	if v, ok := data["name"]; ok {
		r.Name = getStringValue(v)
	}
	if v, ok := data["symbol"]; ok {
		r.Symbol = getStringValue(v)
	}
	if v, ok := data["price_usd"]; ok {
		r.PriceUSD = getStringValue(v)
	}
	if v, ok := data["price_btc"]; ok {
		r.PriceBTC = getStringValue(v)
	}
	if v, ok := data["percent_change_24h"]; ok {
		r.PercentChange24h = getStringValue(v)
	}
	if v, ok := data["last_updated"]; ok {
		r.LastUpdated = getStringValue(v)
	}

	return nil
}

func getStringValue(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	}

	return ""
}

// cryptoCurrency must the full name of the crypto like bitcoin and not btc
func (pc PricingController) getSellPriceKraken(cryptoCurrency CurrencyType, fiatCurrency string) (*CurrencyTickerResponse, error) {
	// URL Request
	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/%s/%s/", APIPathTicker, cryptoCurrency)
	if len(fiatCurrency) > 0 {
		url += fmt.Sprintf("?convert=%s", fiatCurrency)
	}

	log.Printf("url: %s", url)

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
