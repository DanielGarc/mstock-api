module mstock-api/main

go 1.14

replace mstock-api/routes => ./routes

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	mstock-api/routes v0.0.0
)
