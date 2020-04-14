package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// ConnectToDB build the proper connection string
// depending on the enviroment
// if runnnig locally, the sql proxy needs to be running
func ConnectToDB(dbUser string, dbPwd string, dbName string, instanceConnectionName string) (*sql.DB, error) {

	var dbURI string

	mode := os.Getenv("MODE")
	if mode == "dev" {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, "127.0.0.1:3306", dbName)
	} else {
		dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

	}

	log.Println("dburi: ", dbURI)

	// dbPool is the pool of database connections.
	return sql.Open("mysql", dbURI)

}
