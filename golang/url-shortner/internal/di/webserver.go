package di

import (
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/infra/webservers/fiber"
)

func NewWebserver(params ports.WebServerParams) ports.RestWebserver {
	webServer := fiber.New(params)
	return webServer
}
