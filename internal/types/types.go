package types

// DTO структура новости
type NewsDto struct {
	Id         int    `json:"Id"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Categories []int  `json:"Categories"`
}

// Структура новости
type News struct {
	Id      int
	Title   string
	Content string
}

// Структура id новости + id категории
type NewsCategories struct {
	NewsId     int
	CategoryId int
}

// Интерфейс репозитория
type NewsStore interface {
	GetNewsById(id int) (News, error)
	UpdateNews(news News) error
	GetAllNews() ([]News, error)
	GetCategoriesForNews(newsId int) ([]int, error)
	UpdateCategoriesForNews(newsId int, categories []int) error
}
