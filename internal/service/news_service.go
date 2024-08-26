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

func (s NewsService) UpdateNews(news types.News) error {
	// TODO
	return nil
}

func (s NewsService) GetAllNews() ([]types.News, error) {
	// TODO
	return nil, nil
}
