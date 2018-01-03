package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/anothrnick/crypto-historical-data/models"
  "github.com/anothrnick/crypto-historical-data/db"
  "net/http"
)

func GetAllTickers(c *gin.Context) {
	// Get all

	var tickers []models.Ticker
	_tickers := make([]models.TransformedTicker, 0)

	db := db.Database()
    defer db.Close()

	// query for all tickers with last updated
	if err := db.Order("created desc, rank asc").Find(&tickers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error finding tickers."})
		return
	}

	for _, tick := range tickers {
		_tickers = append(
			_tickers, 
			models.TransformedTicker{
						CryptoID: tick.CryptoID,
						Name: tick.Name,
						Symbol: tick.Symbol,
						Rank: tick.Rank,
						PriceUSD: tick.PriceUSD,
						PriceBTC: tick.PriceBTC,
						PercentChange: tick.PercentChange,
						Updated: tick.Updated,
						Created: tick.Created,
					},
				)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"items": _tickers})
}

func GetLatestTickers(c *gin.Context) {
	// Get latest tickers by the last `updated` value

	var tickers []models.Ticker
	var ticker models.Ticker
	_tickers := make([]models.TransformedTicker, 0)

	db := db.Database()
    defer db.Close()
    // get latest updated value
    if err := db.Order("created desc, rank asc").First(&ticker).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error finding ticker."})
		return
	}

	// query for all tickers with last updated
	if err := db.Where("created = ?", ticker.Created).Order("rank asc").Find(&tickers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error finding tickers."})
		return
	}

	for _, tick := range tickers {
		_tickers = append(
			_tickers, 
			models.TransformedTicker{
						CryptoID: tick.CryptoID,
						Name: tick.Name,
						Symbol: tick.Symbol,
						Rank: tick.Rank,
						PriceUSD: tick.PriceUSD,
						PriceBTC: tick.PriceBTC,
						PercentChange: tick.PercentChange,
						Updated: tick.Updated,
						Created: tick.Created,
					},
				)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"items": _tickers})
}

func GetTickerID(c *gin.Context) {
	// Gets the last 12 records for a specific ticker id

	var tickers []models.Ticker
	_tickers := make([]models.TransformedTicker, 0)
	ticker_id := c.Param("id")

	db := db.Database()
    defer db.Close()
	if err := db.Where("crypto_id = ?", ticker_id).Order("created desc").Limit(12).Find(&tickers).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{})
		return
	}

	for _, tick := range tickers {
		_tickers = append(
			_tickers, 
			models.TransformedTicker{
						CryptoID: tick.CryptoID,
						Name: tick.Name,
						Symbol: tick.Symbol,
						Rank: tick.Rank,
						PriceUSD: tick.PriceUSD,
						PriceBTC: tick.PriceBTC,
						PercentChange: tick.PercentChange,
						Updated: tick.Updated,
						Created: tick.Created,
					},
				)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"items": _tickers})
}