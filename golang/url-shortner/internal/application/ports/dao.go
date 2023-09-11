package ports

import "errors"

type Model struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

//go:generate mockgen -source=dao.go -destination ../../infra/persistence/mock/dao_mock.go
type DAO interface {
	FindByShortURL(shortURL string) (*Model, error)
	Save(model Model) error
}

var ShortURLNotFound = errors.New("short url not found")
