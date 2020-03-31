package main

import (
	"mstock-api/packages"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := routes.SetupRouter()
git	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
