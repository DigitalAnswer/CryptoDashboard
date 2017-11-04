package main

import (
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

	var httpClient *http.Client
	httpClient = &http.Client{}
	pc := PricingController{
		client: httpClient,
	}
	result, err := pc.getSellPriceKraken("ETH", settings.currency)
	if err == nil {
		log.Printf("result: %s", *result)
	}
}
