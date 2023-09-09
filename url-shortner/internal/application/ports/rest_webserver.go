package ports

import (
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type WebServerParams struct {
	Port                         int
	RedirectToOriginalURLUseCase UseCase[input.RedirectToOriginalURLUseCaseInput, output.RedirectToOriginalURLUseCaseOutput]
	CreateShortURLUseCase        UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput]
}

type RestWebserver interface {
	ListenAndServe() error
}

type UseCase[input any, output any] interface {
	Execute(input input) (output, error)
}
