package usecases

import (
	"errors"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type CreateShortURLUseCase struct {
	dao       ports.DAO
	encrypter ports.Encrypter
}

func NewCreateShortURLUseCase(
	dao ports.DAO,
	encrypter ports.Encrypter,
) *CreateShortURLUseCase {
	return &CreateShortURLUseCase{
		dao:       dao,
		encrypter: encrypter,
	}
}

func (u *CreateShortURLUseCase) Execute(input input.CreateShortURLUseCaseInput) (output.CreateShortURLUseCaseOutput, error) {
	shortURL, err := u.encrypter.GenerateHash(6)
	if err != nil {
		return output.CreateShortURLUseCaseOutput{}, err
	}
	modelAlreadyExists, _ := u.dao.FindByShortURL(shortURL)
	if modelAlreadyExists != nil {
		return output.CreateShortURLUseCaseOutput{}, errors.New("short url already exists")
	}
	err = u.dao.Save(ports.Model{
		ShortURL:    shortURL,
		OriginalURL: input.OriginalURL,
	})
	if err != nil {
		return output.CreateShortURLUseCaseOutput{}, err
	}
	return output.CreateShortURLUseCaseOutput{
		ShortURL: shortURL,
	}, nil
}
