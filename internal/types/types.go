package types

// Структура новости
type News struct {
	Id         int    `json:"Id"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Categories []int  `json:"Categories"`
}

// Интерфейс репозитория
type NewsStore interface {
	UpdateNews(news News) (News, error)
	GetAllNews() ([]News, error)
}
