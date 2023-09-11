package output

type RedirectToOriginalURLUseCaseOutput struct {
	OriginalURL string
}

type CreateShortURLUseCaseOutput struct {
	ShortURL string `json:"short_url"`
}
