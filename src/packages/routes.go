package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var alphaAPI = "https://www.alphavantage.co/"

var apiKey = os.Getenv("API_KEY")

// SetupRouter list all the api endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/global_quote", globalQuote)

	router.GET("/symbol_search", symbolSearch)

	return router
}

func apiRequest(url string) string {
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

	return string(data[:])
}

// Test ping function to make sure the API is running as expected
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong+"})
}

func globalQuote(c *gin.Context) {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	url := strings.Join([]string{alphaAPI, "query?function=GLOBAL_QUOTE&symbol=", c.Query("symbol"), "&apikey=", apiKey}, "")

	result := apiRequest(url)

	c.JSON(200, gin.H{"message": result})
}

func symbolSearch(c *gin.Context) {
	// https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=BA&apikey=demo
	url := strings.Join([]string{alphaAPI, "query?function=SYMBOL_SEARCH&keywords=", c.Query("keywords"), "&apikey=", apiKey}, "")

	result := apiRequest(url)

	c.JSON(200, gin.H{"message": result})
}
