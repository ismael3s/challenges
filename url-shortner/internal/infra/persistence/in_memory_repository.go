package persistence

import (
	"math/rand"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
)

type inMemoryRepository struct {
	storage map[string]ports.Model
	id      int
}

var _ ports.DAO = &inMemoryRepository{}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		storage: make(map[string]ports.Model),
		id:      rand.Intn(1000),
	}
}

func (r *inMemoryRepository) FindByShortURL(shortURL string) (*ports.Model, error) {
	model, ok := r.storage[shortURL]
	if !ok {
		return nil, ports.ShortURLNotFound
	}
	return &model, nil
}

func (r *inMemoryRepository) Save(model ports.Model) error {
	r.storage[model.ShortURL] = model
	return nil
}
