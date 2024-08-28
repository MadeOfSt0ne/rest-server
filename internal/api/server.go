package api

import (
	"server/internal/db"
	"server/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Запуск сервера
func (s *APIServer) Run() {
	app := fiber.New()

	newsStore := db.NewNewsRepository(s.db)
	newsService := service.NewNewsService(newsStore)
	newsHandler := NewNewsHandler(newsService)
	newsHandler.RegisterRoutes(app)

	logrus.Fatal(app.Listen(s.addr))
}
