package routes

import (
	"os"
)

// We set in this file the structures used to unmarshal the different endpoints from the alphavantage API

var AlphaAPI = "https://www.alphavantage.co/"

var ApiKey = os.Getenv("API_KEY")

type GlobalQuote struct {
	Symbol            string  `json:"01. symbol"`
	Open              float64 `json:"02. open,string"`
	High              float64 `json:"03. high,string"`
	Low               float64 `json:"04. low,string"`
	Price             float64 `json:"05. price,string"`
	Volume            int32   `json:"06. volume,string"`
	LatestTradingDate string  `json:"07. latest trading day"`
	PreviousClose     float64 `json:"08. previous close,string"`
	Change            string  `json:"09. change"`
	ChangePercent     string  `json:"10. change percent"`
	//	XData             interface{} `json:"-"`
}

type SymbolSearch struct {
	Symbol      string  `json:"1. symbol"`
	Name        string  `json:"2. name"`
	Type        string  `json:"3. type"`
	Region      string  `json:"4. region"`
	MarketOpen  string  `json:"5. marketOpen"`
	MarketClose string  `json:"6. marketClose"`
	TimeZone    string  `json:"7. timezone"`
	Currency    string  `json:"8. currency"`
	MatchScore  float64 `json:"9. matchScore,string"`
}
