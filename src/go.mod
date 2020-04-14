module mstock-api/main

go 1.14

replace mstock-api/routes => ./routes

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/joho/godotenv v1.3.0
	mstock-api/routes v0.0.0
)
