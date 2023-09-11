package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

func (s *fiberWebserver) createShortURLHandler(ctx *fiber.Ctx) error {
	var body map[string]string
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, err := s.createShortURLUseCase.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: body["url"],
	})
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	urlWithProtocol := fmt.Sprintf("%s://%s/%s", ctx.Protocol(), ctx.Hostname(), result.ShortURL)
	return ctx.JSON(output.CreateShortURLUseCaseOutput{
		ShortURL: urlWithProtocol,
	})
}
