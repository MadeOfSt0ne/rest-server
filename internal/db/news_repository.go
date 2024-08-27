package db

import (
	"context"
	"server/internal/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsRepository struct {
	dbpool *pgxpool.Pool
}

func NewNewsRepository(pgpool *pgxpool.Pool) *NewsRepository {
	return &NewsRepository{dbpool: pgpool}
}

// Поиск новости по Id
func (r *NewsRepository) GetNewsById(id int) (types.News, error) {
	news := types.News{}
	query := `SELECT * FROM News WHERE Id=$1`
	row := r.dbpool.QueryRow(context.Background(), query, id)
	err := row.Scan(&news.Id, &news.Title, &news.Content)
	return news, err
}

// Обновление новости
func (r *NewsRepository) UpdateNews(news types.News) error {
	query := `UPDATE News SET Title=$1, Content=$2 WHERE Id=$3`
	_, err := r.dbpool.Exec(context.Background(), query, news.Title, news.Content, news.Id)
	return err
}

// Получение списка всех новостей
func (r *NewsRepository) GetAllNews() ([]types.News, error) {
	var news []types.News
	query := `SELECT * FROM News`
	rows, err := r.dbpool.Query(context.Background(), query)
	if err != nil {
		return news, err
	}
	defer rows.Close()

	for rows.Next() {
		n := types.News{}
		err := rows.Scan(&n.Id, &n.Title, &n.Content)
		if err != nil {
			return news, err
		}
		news = append(news, n)
	}
	if err := rows.Err(); err != nil {
		return news, err
	}
	return news, nil
}

// Получение списка категорий для новости по Id
func (r *NewsRepository) GetCategoriesForNews(id int) ([]int, error) {
	categories := []int{}
	query := `SELECT CategoryId FROM NewsCategories WHERE NewsId=$1`
	rows, err := r.dbpool.Query(context.Background(), query, id)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return categories, err
		}
		categories = append(categories, id)
	}
	if err := rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

// Обновление списка категорий для новости по Id
func (r *NewsRepository) UpdateCategoriesForNews(id int, newCats []int) error {
	deleteQuery := `DELETE * FROM NewsCategories WHERE NewsId=$1`
	_, err := r.dbpool.Exec(context.Background(), deleteQuery, id)
	if err != nil {
		return err
	}

	// Обновление бд через цикл выглядит не очень, но другого решения я не нашел
	insertQuery := `UPDATE NewsCategories SET NewsId=$1, CategoryId=$2`
	for _, catId := range newCats {
		_, err := r.dbpool.Exec(context.Background(), insertQuery, id, catId)
		if err != nil {
			return err
		}
	}
	return nil
}
