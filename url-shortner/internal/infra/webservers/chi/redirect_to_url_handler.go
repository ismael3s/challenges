package chi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
)

func (s *chiRestWebServer) redirectToOriginalHandler(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	output, err := s.redirectToOriginalURLUseCase.Execute(input.RedirectToOriginalURLUseCaseInput{
		ShortURL: code,
	})
	fallbackURL := fmt.Sprintf("http://%s/health", r.Host)
	log.Println(fallbackURL)
	if err != nil {
		http.Redirect(w, r, fallbackURL, http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, output.OriginalURL, http.StatusTemporaryRedirect)
}
