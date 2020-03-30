package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var alphaAPI string

var apiKey = os.Getenv("API_KEY")

// Test ping function to make sure the API is running as expected
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func globalQuote(c *gin.Context) {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	client := &http.Client{}
	url := strings.Join([]string{alphaAPI, "query?function=GLOBAL_QUOTE&symbol=", c.Query("symbol"), "&apikey=", apiKey}, "")

	request, err := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": string(data[:])})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/global_quote", globalQuote)

	return router
}

func main() {
	alphaAPI = "https://www.alphavantage.co/"

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := setupRouter()
	router.Run(":" + port)

}
