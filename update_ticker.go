package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/anothrnick/crypto-historical-data/db"
	"github.com/anothrnick/crypto-historical-data/models"
  	"net/http"
	"log"
	"encoding/json"
    "time"
    "strconv"
)

type TickerItem struct {
	ID 					string `json:"id"`
    Name 				string `json:"name"`
    Symbol 				string `json:"symbol"`
    Rank 				string `json:"rank"`
    PriceUSD 			string `json:"price_usd"`
    PriceBTC 			string `json:"price_btc"`
    VolumeUSD24h 		string `json:"24h_volume_usd"`
    MarketCapUSD 		string `json:"market_cap_usd"`
    AvailableSupply 	string `json:"available_supply"` 
    TotalSupply 		string `json:"total_supply"`
    MaxSupply 			string `json:"max_supply"`
    PercentChange1h 	string `json:"percent_change_1h"`
    PercentChange24h 	string `json:"percent_change_24h"`
    PercentChange7d 	string `json:"percent_change_7d"`
    LastUpdated 		string `json:"last_updated"`
}


const TICKER_URL = "https://api.coinmarketcap.com/v1/ticker/"

func main() {
	// get connection to database
	db := db.Database()
    defer db.Close()

    // Build the request
	req, err := http.NewRequest("GET", TICKER_URL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	// send the request via a client
	// sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var items []TickerItem

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		log.Println(err)
	}

    created := time.Now().Unix()
	for _, ticker := range items {
		// convert String unix timestamp to time
		timeInt, err := strconv.ParseInt(ticker.LastUpdated, 10, 64)
	    if err != nil {
	        panic(err)
	    }
	    updated := time.Unix(timeInt, 0)

		newTick := models.Ticker{
						CryptoID: ticker.ID,
						Name: ticker.Name,
						Symbol: ticker.Symbol,
						Rank: ticker.Rank,
						PriceUSD: ticker.PriceUSD,
						PriceBTC: ticker.PriceBTC,
						PercentChange: ticker.PercentChange1h,
						Updated: updated,
						Created: created,
					}

		db.Save(&newTick)
	}
}