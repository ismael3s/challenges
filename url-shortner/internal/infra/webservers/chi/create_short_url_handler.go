package chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/output"
)

func (s *chiRestWebServer) createShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := s.createShortURLUseCase.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: body["url"],
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var protocol string
	if r.TLS == nil {
		protocol = "http"
	} else {
		protocol = "https"
	}
	urlWithProtocol := fmt.Sprintf("%s://%s/%s", protocol, r.Host, result.ShortURL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output.CreateShortURLUseCaseOutput{
		ShortURL: urlWithProtocol,
	})
}
