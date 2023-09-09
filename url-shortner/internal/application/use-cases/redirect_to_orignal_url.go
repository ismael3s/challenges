package usecases

import (
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type RedirectToOriginalURLUseCase struct {
	dao ports.DAO
}

func NewRedirectToOriginalURLUseCase(
	dao ports.DAO,
) *RedirectToOriginalURLUseCase {
	return &RedirectToOriginalURLUseCase{
		dao: dao,
	}
}

func (u *RedirectToOriginalURLUseCase) Execute(input input.RedirectToOriginalURLUseCaseInput) (output.RedirectToOriginalURLUseCaseOutput, error) {
	model, err := u.dao.FindByShortURL(input.ShortURL)
	if err != nil {
		return output.RedirectToOriginalURLUseCaseOutput{}, err
	}
	return output.RedirectToOriginalURLUseCaseOutput{
		OriginalURL: model.OriginalURL,
	}, nil
}
