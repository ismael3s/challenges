package encrypter

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
)

type encrypter struct{}

func New() ports.Encrypter {
	return &encrypter{}
}

func (e *encrypter) GenerateHash(maxLength int) (string, error) {
	alpha := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-!@#$%")
	sb := strings.Builder{}
	for i := 0; i < 6; i++ {
		position, err := rand.Int(rand.Reader, big.NewInt(int64(len(alpha))))
		if err != nil {
			return "", err
		}
		sb.WriteRune(alpha[position.Int64()])
	}
	shortURL := sb.String()
	return shortURL, nil
}
