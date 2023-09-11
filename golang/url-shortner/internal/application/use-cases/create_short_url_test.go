package usecases

import (
	"errors"
	"testing"

	"github.com/ismael3s/challenges/url-shortner/internal/application/ports"
	"github.com/ismael3s/challenges/url-shortner/internal/application/use-cases/input"
	"github.com/ismael3s/challenges/url-shortner/internal/infra/encrypter"
	mockEncrypter "github.com/ismael3s/challenges/url-shortner/internal/infra/encrypter/mock"
	"github.com/ismael3s/challenges/url-shortner/internal/infra/persistence"
	mockDAO "github.com/ismael3s/challenges/url-shortner/internal/infra/persistence/mock"
	"go.uber.org/mock/gomock"
)

func TestShouldBeAbleToGenerateAShortURL(t *testing.T) {
	// Arrange
	sut := NewCreateShortURLUseCase(persistence.NewInMemoryRepository(), encrypter.New())
	// Act
	result, err := sut.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: "https://www.google.com",
	})
	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if result.ShortURL == "" {
		t.Errorf("Expected short url to be generated but got empty string")
	}
}

func TestGivenAURLToBeShortnedAndTheGeneratedHashHasAlreadyGeneratedBeforeThenShouldReturnAnError(t *testing.T) {
	// Arrange
	dao := persistence.NewInMemoryRepository()
	dao.Save(ports.Model{
		ShortURL:    "abc123",
		OriginalURL: "https://www.google.com",
	})
	ctrl := gomock.NewController(t)
	encrypter := mockEncrypter.NewMockEncrypter(ctrl)
	sut := NewCreateShortURLUseCase(dao, encrypter)
	encrypter.EXPECT().GenerateHash(6).Return("abc123", nil)
	// Act
	result, err := sut.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: "https://www.google.com",
	})
	// Assert
	if err == nil {
		t.Error("Want error but succeed")
	}
	if result.ShortURL != "" {
		t.Errorf("Expected short url to be empty but got %s", result.ShortURL)
	}
}

func TestWhenFailsToGenerateHashShouldForwardTheError(t *testing.T) {
	// Arrange
	expectedError := errors.New("some error")
	dao := persistence.NewInMemoryRepository()
	ctrl := gomock.NewController(t)
	encrypter := mockEncrypter.NewMockEncrypter(ctrl)
	sut := NewCreateShortURLUseCase(dao, encrypter)

	encrypter.EXPECT().GenerateHash(6).Return("", expectedError)
	// Act
	result, err := sut.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: "https://www.google.com",
	})
	// Assert
	if err == nil {
		t.Error("Want error but succeed")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error to be %s but got %s", expectedError.Error(), err.Error())
	}
	if result.ShortURL != "" {
		t.Errorf("Expected short url to be empty but got %s", result.ShortURL)
	}
}

func TestWhenFailsToPersisteGeneratedShortURL_ShouldForwardError(t *testing.T) {
	// Arrange
	expectedError := errors.New("some error on the database")
	ctrl := gomock.NewController(t)
	encrypter := mockEncrypter.NewMockEncrypter(ctrl)
	dao := mockDAO.NewMockDAO(ctrl)
	sut := NewCreateShortURLUseCase(dao, encrypter)
	encrypter.EXPECT().GenerateHash(6).Return("abc123", nil)
	dao.EXPECT().
		FindByShortURL(gomock.Any()).
		Return(nil, ports.ShortURLNotFound)
	dao.EXPECT().
		Save(gomock.Any()).
		Return(expectedError)
	// Act
	result, err := sut.Execute(input.CreateShortURLUseCaseInput{
		OriginalURL: "https://www.google.com",
	})
	// Assert
	if err == nil {
		t.Error("Want error but succeed")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error to be %s but got %s", expectedError.Error(), err.Error())
	}
	if result.ShortURL != "" {
		t.Errorf("Expected short url to be empty but got %s", result.ShortURL)
	}
}
