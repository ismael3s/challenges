package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type fiberWebserver struct {
	port                         int
	app                          *fiber.App
	redirectToOriginalURLUseCase ports.UseCase[input.RedirectToOriginalURLUseCaseInput, output.RedirectToOriginalURLUseCaseOutput]
	createShortURLUseCase        ports.UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput]
}

func New(params ports.WebServerParams) *fiberWebserver {
	app := fiber.New()
	server := &fiberWebserver{
		port:                         params.Port,
		app:                          app,
		redirectToOriginalURLUseCase: params.RedirectToOriginalURLUseCase,
		createShortURLUseCase:        params.CreateShortURLUseCase,
	}
	server.registerRoutes()
	return server
}

func (s *fiberWebserver) registerRoutes() {
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK, Fiber")
	})

	s.app.Get("/:code", s.redirectToOriginalHandler)
	s.app.Post("/", s.createShortURLHandler)
}

func (s *fiberWebserver) ListenAndServe() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}
