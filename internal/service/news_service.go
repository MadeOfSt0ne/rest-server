package service

import (
	"rest-server/internal/types"
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
		return news, err
	}
	if news.Title != "" {
		old.Title = news.Title
	}
	if news.Content != "" {
		old.Content = news.Content
	}
	if len(news.Categories) != 0 {
		s.store.UpdateCategoriesForNews(news.Id, news.Categories)
	}
	err = s.store.UpdateNews(old)
	if err != nil {
		return news, err
	}

	fromDB, err := s.store.GetNewsById(news.Id)
	if err != nil {
		return news, err
	}
	categories := s.store.GetCategoriesForNews(news.Id)
	updated := types.NewsDto{
		Id:         fromDB.Id,
		Title:      fromDB.Title,
		Content:    fromDB.Content,
		Categories: categories,
	}
	return updated, nil
}

// Получение списка всех новостей
func (s NewsService) GetAllNews() ([]types.NewsDto, error) {
	newsWithCategories := []types.NewsDto{}
	news, err := s.store.GetAllNews()
	if err != nil {
		return newsWithCategories, err
	}

	for _, val := range news {
		categories := s.store.GetCategoriesForNews(val.Id)
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
