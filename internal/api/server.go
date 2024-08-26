package api

import (
	"database/sql"
	"rest-server/internal/db"
	"rest-server/internal/service"

	"github.com/gofiber/fiber/v3"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {
	app := fiber.New()

	newsStore := db.NewNewsRepository(s.db)
	newsService := service.NewNewsService(newsStore)
	newsHandler := NewNewsHandler(newsService)
	newsHandler.RegisterRoutes(app)

	app.Listen(s.addr)
}
