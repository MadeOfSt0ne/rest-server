package api

import (
	"encoding/json"
	"server/internal/service"
	"server/internal/types"

	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/sirupsen/logrus"
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
	lim := c.Query("limit")

	limit, err := strconv.Atoi(lim)
	if err != nil && err != strconv.ErrSyntax {
		log.Infof("Неверный формат limit: %v", err)
		return fiber.NewError(400, "Неверный формат количества страниц")
	}

	offs := c.Query("offset")
	offset, err := strconv.Atoi(offs)
	if err != nil && err != strconv.ErrSyntax {
		log.Infof("Неверный формат offset: %v", err)
		return fiber.NewError(400, "Неверный формат офсета")
	}

	news, err := h.srv.GetAllNews(limit, offset)
	if err != nil {
		logrus.Errorf("Ошибка бд: %v", err)
		return fiber.NewError(500, "Что-то пошло не так")
	}

	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"News":    news,
	})
}

// Обновление новости
func (h *Handler) handleUpdateNews(c fiber.Ctx) error {
	idString := c.Params("Id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Infof("Неверный формат id: %v", err)
		return fiber.NewError(400, "Неверный формат id")
	}

	var news types.NewsDto
	if err = json.Unmarshal(c.Body(), &news); err != nil {
		logrus.WithFields(logrus.Fields{
			"news":  news,
			"error": err,
		}).Debug("Ошибка при анмаршале json")
		return fiber.NewError(400, "Неверный формат данных")
	}
	if id != news.Id {
		logrus.Info("Попытка изменить id")
		return fiber.NewError(400, "Нельзя изменить Id новости")
	}

	updated, err := h.srv.UpdateNews(news)
	if err != nil {
		logrus.Errorf("Ошибка бд: %v", err)
		return fiber.NewError(500, "Что-то пошло не так")
	}
	return c.Status(200).JSON(updated)
}
