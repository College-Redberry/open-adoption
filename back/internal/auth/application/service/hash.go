package service

type HashService interface {
	Hash(text string) (string, error)
	Verify(text string, hashedText string) (bool, error)
}
