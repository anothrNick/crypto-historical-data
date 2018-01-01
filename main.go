package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/anothrnick/crypto-historical-data/db"
	"github.com/anothrnick/crypto-historical-data/controllers"
	"github.com/anothrnick/crypto-historical-data/models"
	"github.com/anothrnick/crypto-historical-data/middleware"
  	"net/http"
)

func NotImplemented(c * gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

func main() {
	//Migrate schema
	db := db.Database()
	db.AutoMigrate(&models.Ticker{})

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api")
	{
		//v1.GET("/", controllers.GetAllTickers)
		v1.GET("/ticker", controllers.GetLatestTickers)
		v1.GET("/ticker/:id", controllers.GetTickerID)
	}

	router.Run(":5001")
}