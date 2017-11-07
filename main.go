package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB

var settings = Settings{
	currency: "EUR",
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
	result, err := pc.getSellPriceKraken(CurrencyTypeBTC, settings.currency)
	if err == nil {
		log.Printf("result: %s", *result)
	}

	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", "root:admin@tcp(localhost:6603)/crypto_dashboard")
	if err != nil {
		panic(err.Error())
	}

	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
