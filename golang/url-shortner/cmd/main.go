package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/di"
	"github.com/ismael3s/challenges/url-shortner/internal/infra/encrypter"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "4000"
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("PORT env var must be a number")
	}
	repository := di.NewRepository()
	encrypter := encrypter.New()
	useCase := di.NewRedirectToOriginalURLUseCase(repository)
	createShortURLUseCase := di.NewCreateShortURLUseCaseDecorated(repository, encrypter)
	server := di.NewWebserver(ports.WebServerParams{
		Port:                         port,
		RedirectToOriginalURLUseCase: useCase,
		CreateShortURLUseCase:        createShortURLUseCase,
	})

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
