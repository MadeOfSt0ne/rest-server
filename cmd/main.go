package main

import (
	"context"
	"fmt"
	"os"
	"rest-server/internal/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool := connectDB()

	server := api.NewAPIServer(os.Getenv("SERVER_PORT"), dbpool)
	server.Run()
	defer dbpool.Close()
}

func connectDB() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	createTable(dbpool)
	return dbpool
}

func createTable(dbpool *pgxpool.Pool) {
	create := `
	    CREATE TABLE IF NOT EXISTS News(
		    Id BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
			Title VARCHAR(256) NOT NULL,
		    Content VARCHAR NOT NULL,
			PRIMARY KEY (Id)
		);
		CREATE TABLE IF NOT EXISTS NewsCategories(
		    NewsId BIGINT NOT NULL,
			CategoryId BIGINT NOT NULL,
            PRIMARY KEY (NewsId, CategoryId)
		);`
	_, err := dbpool.Exec(context.Background(), create)
	if err != nil {

	}
}
