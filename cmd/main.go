package main

import (
	"context"
	"fmt"
	"os"
	"server/internal/api"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()
	dbpool := connectDB()

	server := api.NewAPIServer(os.Getenv("SERVER_PORT"), dbpool)
	server.Run()
	defer dbpool.Close()
}

// Загрузка env переменных
func loadEnv() {
	logrus.Info("Загрузка env файла")
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Не удалось загрузить env файл")
	}
}

// Подключение к бд (connection pool)
func connectDB() *pgxpool.Pool {
	logrus.Info("Соединение с бд")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		logrus.Fatalf("Ошибка при подключении к бд: %v", err)
	}
	defer dbpool.Close()

	createTable(dbpool)
	return dbpool
}

// Создание таблиц в бд
func createTable(dbpool *pgxpool.Pool) {
	logrus.Info("Создание таблиц в бд")
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
		logrus.Debugf("Ошибка при создании таблиц: %v", err)
	}
}
