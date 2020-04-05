package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var alphaAPI = "https://www.alphavantage.co/"

var apiKey = os.Getenv("API_KEY")

// We need to make this a "generic" Object since sometimes it might get different parameters.
type Symbol struct {
	Name              string      `json:"01. symbol"`
	Open              float64     `json:"02. open,string"`
	High              float64     `json:"03. high,string"`
	Low               float64     `json:"04. low,string"`
	Price             float64     `json:"05. price,string"`
	Volume            int32       `json:"06. volume,string"`
	LatestTradingDate string      `json:"07. latest trading day"`
	PreviousClose     float64     `json:"08. previous close,string"`
	Change            string      `json:"09. change"`
	ChangePercent     string      `json:"10. change percent"`
	XData             interface{} `json:"-"`
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

	c.JSON(http.StatusOK, gin.H{"Name": symbol.Name, "Open": symbol.Open, "High": symbol.High, "Low": symbol.Low, "Price": symbol.Price, "Volume": symbol.Volume})
}

func symbolSearch(c *gin.Context) {
	// https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=BA&apikey=demo
	url := strings.Join([]string{alphaAPI, "query?function=SYMBOL_SEARCH&keywords=", c.Query("keywords"), "&apikey=", apiKey}, "")

	result := apiRequest(url)

	//c.JSON(200, gin.H{"message": result})
	c.String(http.StatusOK, string(result))
}
