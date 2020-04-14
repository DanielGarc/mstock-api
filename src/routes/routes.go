package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRouter list all the api endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/global_quote", globalQuoteEndpoint)

	router.GET("/symbol_search", symbolSearchEndpoint)

	return router
}

// To be able to know the structure of the JSON we are trying to read
// These type of "tools" should be added to a separate script or maybe to an option in the CLI
func decryptUnknownJson(data []byte) {
	var f interface{}

	if err := json.Unmarshal([]byte(data), &f); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", f)
}

func apiRequest(url string) []byte {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//	decryptUnknownJson(data)

	return data
}

// Test ping function to make sure the API is running as expected
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong++"})
}

func globalQuoteEndpoint(c *gin.Context) {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	url := strings.Join([]string{AlphaAPI, "query?function=GLOBAL_QUOTE&symbol=", c.Query("symbol"), "&apikey=", ApiKey}, "")

	quote := globalQuote(url)

	//c.String(http.StatusOK, string(result))
	c.JSON(http.StatusOK, gin.H{"Symbol": quote.Symbol, "Open": quote.Open, "High": quote.High, "Low": quote.Low, "Price": quote.Price, "Volume": quote.Volume})
}

func globalQuote(url string) GlobalQuote {
	result := apiRequest(url)

	var s map[string]GlobalQuote
	if err := json.Unmarshal([]byte(result), &s); err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", s["Global Quote"])

	quote := s["Global Quote"]

	return quote
}

func symbolSearchEndpoint(c *gin.Context) {
	// https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=BA&apikey=demo
	url := strings.Join([]string{AlphaAPI, "query?function=SYMBOL_SEARCH&keywords=", c.Query("keywords"), "&apikey=", ApiKey}, "")

	matches := symbolSearch(url)

	c.JSON(200, gin.H{"bestMatches": matches})
}

func symbolSearch(url string) []SymbolSearch {
	result := apiRequest(url)

	var s map[string][]SymbolSearch
	if err := json.Unmarshal([]byte(result), &s); err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", s["bestMatches"])

	matches := s["bestMatches"]

	return matches
}
