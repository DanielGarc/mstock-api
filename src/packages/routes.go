package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var alpha_API = "https://www.alphavantage.co/"

var api_key = os.Getenv("API_KEY")

// Test ping function to make sure the API is running as expected
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func globalQuote(c *gin.Context) {
	//	https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo
	client := &http.Client{}
	url := strings.Join([]string{alpha_API, "query?function=GLOBAL_QUOTE&symbol=", c.Query("symbol"), "&apikey=", api_key}, "")

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

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/global_quote", global_quote)

	return router
}
