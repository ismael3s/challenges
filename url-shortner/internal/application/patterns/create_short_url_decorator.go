package patterns

import (
	"log"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type CreateShortURLUseCaseDecorator struct {
	useCase    ports.UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput]
	repository ports.DAO
	encrypter  ports.Encrypter
}

func NewCreateShortURLUseCaseDecorator(
	useCase ports.UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput],
	repository ports.DAO,
	encrypter ports.Encrypter,
) *CreateShortURLUseCaseDecorator {
	return &CreateShortURLUseCaseDecorator{
		useCase:    useCase,
		repository: repository,
		encrypter:  encrypter,
	}
}

func (u *CreateShortURLUseCaseDecorator) Execute(input input.CreateShortURLUseCaseInput) (output.CreateShortURLUseCaseOutput, error) {
	retriesCount := 0
	maxRetries := 5
begin:
	result, err := u.useCase.Execute(input)
	if err != nil && retriesCount < maxRetries {
		log.Println("Retrying to create short url")
		retriesCount++
		goto begin
	}
	return result, err
}
