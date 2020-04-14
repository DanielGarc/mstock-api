package main

import (
	"log"
	"mstock-api/main/database"
	"mstock-api/routes"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	//Load env file for local use

	if os.Getenv("MODE") == "dev" {

		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Database env variables
	var (
		dbUser                 = os.Getenv("DB_USER")
		dbPwd                  = os.Getenv("DB_PWD")
		instanceConnectionName = os.Getenv("INSTANCE_CONNETION_NAME")
		dbName                 = "temp"
	)

	db, err := database.ConnectToDB(dbUser, dbPwd, dbName, instanceConnectionName)

	err = db.Ping()
	if err != nil {
		log.Fatal("sql ping error: ", err)
	}

	router := routes.SetupRouter()
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
