package main

import (
	"log"
	"mstock-api/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := routes.New("RCL")
	log.Printf("%+v", s)

	router := routes.SetupRouter()
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
