package di

import (
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/infra/persistence"
)

func NewRepository() ports.DAO {
	return persistence.NewInMemoryRepository()
}
