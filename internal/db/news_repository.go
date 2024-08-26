package db

import (
	"database/sql"
	"rest-server/internal/types"
)

type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) UpdateNews(news types.News) (types.News, error) {
	// TODO
	return nil, nil
}

func (r *NewsRepository) GetAllNews() ([]types.News, error) {
	// TODO
	return nil, nil
}
