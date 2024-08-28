package service

import (
	"fmt"
	"server/internal/types"

	"github.com/sirupsen/logrus"
)

type NewsService struct {
	store types.NewsStore
}

func NewNewsService(store types.NewsStore) NewsService {
	return NewsService{store: store}
}

// Обновление новости (если поле пустое - оно не обновляется)
func (s NewsService) UpdateNews(news types.NewsDto) (types.NewsDto, error) {
	old, err := s.store.GetNewsById(news.Id)
	if err != nil {
		logrus.Errorf("Ошибка бд: %v", err)
		return news, fmt.Errorf("что-то пошло не так")
	}
	if news.Title != "" {
		old.Title = news.Title
	}
	if news.Content != "" {
		old.Content = news.Content
	}
	if len(news.Categories) != 0 {
		err := s.store.UpdateCategoriesForNews(news.Id, news.Categories)
		logrus.Errorf("Ошибка при обновлении категорий в бд: %v", err)
	}
	err = s.store.UpdateNews(old)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении новости в бд: %v", err)
		return news, fmt.Errorf("что-то пошло не так")
	}

	fromDB, err := s.store.GetNewsById(news.Id)
	if err != nil {
		logrus.Errorf("Ошибка при получении новости из бд: %v", err)
		return news, fmt.Errorf("что-то пошло не так")
	}
	categories, err := s.store.GetCategoriesForNews(news.Id)
	if err != nil {
		logrus.Errorf("Ошибка при получении категорий новости из бд: %v", err)
		return news, fmt.Errorf("что-то пошло не так")
	}
	updated := types.NewsDto{
		Id:         fromDB.Id,
		Title:      fromDB.Title,
		Content:    fromDB.Content,
		Categories: categories,
	}
	return updated, nil
}

// Получение списка всех новостей
func (s NewsService) GetAllNews(limit, offset int) ([]types.NewsDto, error) {
	if limit == 0 {
		limit = 10
	}

	newsWithCategories := []types.NewsDto{}
	news, err := s.store.GetAllNews(limit, offset)
	if err != nil {
		logrus.Errorf("Ошибка при получении списка новостей из бд: %v", err)
		return newsWithCategories, fmt.Errorf("что-то пошло не так")
	}

	for _, val := range news {
		categories, err := s.store.GetCategoriesForNews(val.Id)
		if err != nil {
			logrus.Errorf("Ошибка при получении списка категорий из бд: %v", err)
		}
		newsWithCategories = append(newsWithCategories,
			types.NewsDto{
				Id:         val.Id,
				Title:      val.Title,
				Content:    val.Content,
				Categories: categories,
			})
	}
	return newsWithCategories, nil
}
