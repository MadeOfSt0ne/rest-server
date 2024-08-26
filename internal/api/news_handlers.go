package api

import (
	"rest-server/internal/service"
	"rest-server/internal/types"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	srv service.NewsService
}

func NewNewsHandler(srv service.NewsService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/edit/:Id", h.handleUpdateNews)
	app.Get("/list", h.handleGetNews)
}

// Получение списка новостей
func (h *Handler) handleGetNews(c fiber.Ctx) error {
	news, err := h.srv.GetAllNews()
	if err != nil {
		return fiber.NewError(500, "Что-то пошло не так")
	}

	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"News":    news,
	})
}

// Обновление новости
func (h *Handler) handleUpdateNews(c fiber.Ctx) error {
	idString := c.FormValue("Id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return fiber.NewError(400, "Неверный формат id")
	}

	var body types.News
	if err = c.BodyParser(body); err != nil {
		return fiber.NewError(400, "Неверный формат данных")
	}
	if id != body.Id {
		return fiber.NewError(400, "Нельзя изменить Id новости")
	}

	return c.Status(200).JSON(updated)
}
