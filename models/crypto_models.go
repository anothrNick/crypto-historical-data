package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Ticker struct {
	gorm.Model
	CryptoID		string 		`json:"id"`
	Name 			string 		`json:"name"`
	Symbol 			string 		`json:"symbol"`
	Rank			int 		`json:"rank"`
	PriceUSD		string 		`json:"price_usd"`
	PriceBTC		string 		`json:"price_btc"`
	PercentChange	string 		`json:"percent_change_1h"`
	Updated 		time.Time 	`json:"updated"`
	Created 		int64 		`json:"created"`
}

type TransformedTicker struct {
	CryptoID		string 		`json:"id"`
	Name 			string 		`json:"name"`
	Symbol 			string 		`json:"symbol"`
	Rank			int 		`json:"rank"`
	PriceUSD		string 		`json:"price_usd"`
	PriceBTC		string 		`json:"price_btc"`
	PercentChange	string 		`json:"percent_change_1h"`
	Updated 		time.Time 	`json:"updated"`
	Created 		int64 		`json:"created"`
}
