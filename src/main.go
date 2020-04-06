package main

import (
	"mstock-api/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRouter()
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
