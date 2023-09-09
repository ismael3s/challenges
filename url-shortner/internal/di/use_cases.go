package di

import (
	"github.com/ismael3s/challenges/url-shortner/internal/application/patterns"
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	usecases "github.com/ismael3s/challenges/url-shortner/internal/application/use-cases"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

func NewRedirectToOriginalURLUseCase(
	repository ports.DAO,
) *usecases.RedirectToOriginalURLUseCase {
	return usecases.NewRedirectToOriginalURLUseCase(repository)
}

func NewCreateShortURLUseCase(
	repository ports.DAO,
	encrypter ports.Encrypter,
) *usecases.CreateShortURLUseCase {
	return usecases.NewCreateShortURLUseCase(repository, encrypter)
}

func NewCreateShortURLUseCaseDecorated(
	repository ports.DAO,
	encrypter ports.Encrypter,
) ports.UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput] {
	useCase := usecases.NewCreateShortURLUseCase(repository, encrypter)
	return patterns.NewCreateShortURLUseCaseDecorator(useCase, repository, encrypter)
}
