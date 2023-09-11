package fiber

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
)

func (s *fiberWebserver) redirectToOriginalHandler(ctx *fiber.Ctx) error {
	shortURL := ctx.Params("code")
	output, err := s.redirectToOriginalURLUseCase.Execute(input.RedirectToOriginalURLUseCaseInput{
		ShortURL: shortURL,
	})
	fallbackURL := fmt.Sprintf("http://%s/health", ctx.Hostname())
	if err != nil {
		return ctx.Redirect(fallbackURL, http.StatusTemporaryRedirect)

	}
	return ctx.Redirect(output.OriginalURL, http.StatusTemporaryRedirect)
}
