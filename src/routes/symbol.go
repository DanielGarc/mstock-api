package routes

import (
	"strings"
	"time"
)

type Symbol struct {
	Code       string
	Name       string
	Type       string
	Region     string
	Currency   string
	Price      float64
	Volume     int32
	TodaysHigh Peak
	TodaysLow  Peak
}

type Peak struct {
	Value  float64
	Date   time.Time
	IsHigh bool
}

func New(code string) Symbol {
	quote := getSymbolNumbers(code)
	match := getSymbolInfo(code)

	todaysHigh := Peak{quote.High, time.Now(), true}
	todaysLow := Peak{quote.Low, time.Now(), false}

	s := Symbol{code, match.Name, match.Type, match.Region, match.Currency, quote.Price, quote.Volume, todaysHigh, todaysLow}

	return s
}

var getSymbolNumbers = func(code string) GlobalQuote {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	url := strings.Join([]string{AlphaAPI, "query?function=GLOBAL_QUOTE&symbol=", code, "&apikey=", ApiKey}, "")

	quote := globalQuote(url)

	return quote
}

var getSymbolInfo = func(code string) SymbolSearch {
	// https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=BA&apikey=demo
	url := strings.Join([]string{AlphaAPI, "query?function=SYMBOL_SEARCH&keywords=", code, "&apikey=", ApiKey}, "")

	matches := symbolSearch(url)

	var match SymbolSearch

	for _, v := range matches {
		if v.MatchScore == 1 {
			match = v
			break
		}
	}

	return match
}
