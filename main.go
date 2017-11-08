package main

import (
	"database/sql"
	"fmt"
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

	// http.ListenAndServe(":8080", nil)

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

	listAllCrypto()

	_, err = db.Exec("INSERT INTO btc_prices(last_update, currency_code, price_usd, price_btc, PercentChange24h) VALUES(?, ?, ?, ?, ?)", result.LastUpdated, result.Symbol, result.PriceUSD, result.PriceBTC, result.PercentChange24h)
	if err != nil {
		log.Println(err)
		return
	}

	listAllData()
}

func listAllCrypto() {

	// Grab from the database
	var databaseCode string
	var databaseName string

	err := db.QueryRow("SELECT code, name FROM currency_info WHERE code='BTC'").Scan(&databaseCode, &databaseName)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("result db: %s-%s", databaseCode, databaseName)
}

func listAllData() {

	rows, err := db.Query("SELECT COUNT(*) as count FROM btc_prices")
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}

	fmt.Println("Total count:", count)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
