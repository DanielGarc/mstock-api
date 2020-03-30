package main

import (
	"mstock-api/packages"
)

func main() {
	router := routes.SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
