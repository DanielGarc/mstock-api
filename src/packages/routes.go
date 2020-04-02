package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var alphaAPI = "https://www.alphavantage.co/"

var apiKey = os.Getenv("API_KEY")

type Symbol struct {
	Name              string `json:"01. symbol"`
	Open              string `json:"02. open"`
	High              string `json:"03. high"`
	Low               string `json:"04. low"`
	Price             string `json:"05. price"`
	Volume            string `json:"06. volume"`
	LatestTradingDate string `json:"07. latest trading day"`
	PreviousClose     string `json:"08. previous close"`
	Change            string `json:"09. change"`
	ChangePercent     string `json:"10. change percent"`
}

// SetupRouter list all the api endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/global_quote", globalQuote)

	router.GET("/symbol_search", symbolSearch)

	return router
}

// To be able to know the structure of hte JSON we are trying to read
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
	c.JSON(200, gin.H{"message": "pong+"})
}

func globalQuote(c *gin.Context) {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	url := strings.Join([]string{alphaAPI, "query?function=GLOBAL_QUOTE&symbol=", c.Query("symbol"), "&apikey=", apiKey}, "")

	result := apiRequest(url)

	var s map[string]Symbol
	if err := json.Unmarshal([]byte(result), &s); err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", s["Global Quote"])

	symbol := s["Global Quote"]

	//c.String(http.StatusOK, string(result))

	open, _ := strconv.ParseFloat(symbol.Open, 64)
	high, _ := strconv.ParseFloat(symbol.High, 64)
	low, _ := strconv.ParseFloat(symbol.Low, 64)
	price, _ := strconv.ParseFloat(symbol.Price, 64)
	volume, _ := strconv.ParseInt(symbol.Volume, 10, 32)

	c.JSON(http.StatusOK, gin.H{"Name": symbol.Name, "Open": open, "High": high, "Low": low, "Price": price, "Volume": volume})
}

func symbolSearch(c *gin.Context) {
	// https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=BA&apikey=demo
	url := strings.Join([]string{alphaAPI, "query?function=SYMBOL_SEARCH&keywords=", c.Query("keywords"), "&apikey=", apiKey}, "")

	result := apiRequest(url)

	//c.JSON(200, gin.H{"message": result})
	c.String(http.StatusOK, string(result))
}
