package ports

//go:generate mockgen -source=encrypter.go -destination ../../infra/encrypter/mock/encrypter_mock.go
type Encrypter interface {
	GenerateHash(maxLength int) (string, error)
}
