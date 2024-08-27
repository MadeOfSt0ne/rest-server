package api

import (
	"encoding/json"
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

	var news types.NewsDto
	if err = json.Unmarshal(c.Body(), &news); err != nil {
		return fiber.NewError(400, "Неверный формат данных")
	}
	if id != news.Id {
		return fiber.NewError(400, "Нельзя изменить Id новости")
	}

	updated, err := h.srv.UpdateNews(news)
	if err != nil {
		return fiber.NewError(500, "Что-то пошло не так")
	}
	return c.Status(200).JSON(updated)
}
