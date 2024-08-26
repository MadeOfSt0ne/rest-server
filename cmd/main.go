package main

import (
	"database/sql"
	"rest-server/internal/api"
)

func main() {
	loadEnv()
	db := connectDB()

	server := api.NewAPIServer(":8080", db)
	server.Run()
	defer db.Close()
}

func loadEnv() {

}

func connectDB() *sql.DB {

}
