package chi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

type chiRestWebServer struct {
	router                       chi.Router
	port                         int
	createShortURLUseCase        ports.UseCase[input.CreateShortURLUseCaseInput, output.CreateShortURLUseCaseOutput]
	redirectToOriginalURLUseCase ports.UseCase[input.RedirectToOriginalURLUseCaseInput, output.RedirectToOriginalURLUseCaseOutput]
}

func New(params ports.WebServerParams) *chiRestWebServer {
	router := chi.NewRouter()
	server := &chiRestWebServer{
		router:                       router,
		port:                         params.Port,
		redirectToOriginalURLUseCase: params.RedirectToOriginalURLUseCase,
		createShortURLUseCase:        params.CreateShortURLUseCase,
	}
	server.registerRoutes()
	return server
}

func (s *chiRestWebServer) registerRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK CHI"))
	})

	s.router.Get("/{code}", s.redirectToOriginalHandler)
	s.router.Post("/", s.createShortURLHandler)
}

func (s *chiRestWebServer) GetHandler() http.Handler {
	return s.router
}

func (s *chiRestWebServer) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
