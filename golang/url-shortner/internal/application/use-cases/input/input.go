package input

type RedirectToOriginalURLUseCaseInput struct {
	ShortURL string
}

type CreateShortURLUseCaseInput struct {
	OriginalURL string
	Host        string
}
