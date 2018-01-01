## Cryptocurrency Historical Data
This application collects data from [Coin Market Cap](https://coinmarketcap.com/api/) in order to provide a RESTful interface to historical information regarding each ticker.

### Start server on localhost
```bash
$ go run main.go
```

### Start with Docker
```bash
# build api
$ make build

# start container
$ make up
```

#### TODO
* Positions
* Graph market history

#### API Endpoints
* [Get latest tickers - GET /api/ticker](#get-apiticker)   
* [Get history for a specific crypto - GET /api/ticker/:id](#get-apitickerid)             

#### GET /api/ticker
```bash
$ curl -s /api/ticker
{
  "items": [
    {
      "id": "bitcoin",
      "name": "Bitcoin",
      "symbol": "BTC",
      "rank": "1",
      "price_usd": "13488.9",
      "price_btc": "1.0",
      "percent_change_1h": "0.77",
      "updated": "2018-01-01T15:04:21-05:00",
      "created": 1514837193
    },
    {
      "id": "litecoin",
      "name": "Litecoin",
      "symbol": "LTC",
      "rank": "6",
      "price_usd": "226.556",
      "price_btc": "0.0170479",
      "percent_change_1h": "-0.19",
      "updated": "2018-01-01T15:04:01-05:00",
      "created": 1514837193
    },
    ...
  ]
}
```

#### GET /api/ticker/:id
```bash
$ curl -s /api/ticker/bitcoin
{
  "items": [
    {
      "id": "bitcoin",
      "name": "Bitcoin",
      "symbol": "BTC",
      "rank": "1",
      "price_usd": "13488.9",
      "price_btc": "1.0",
      "percent_change_1h": "0.77",
      "updated": "2018-01-01T15:04:21-05:00",
      "created": 1514837193
    },
    {
      "id": "bitcoin",
      "name": "Bitcoin",
      "symbol": "BTC",
      "rank": "1",
      "price_usd": "13495.5",
      "price_btc": "1.0",
      "percent_change_1h": "0.75",
      "updated": "2018-01-01T14:59:20-05:00",
      "created": 1514837040
    }
  ]
}
```
