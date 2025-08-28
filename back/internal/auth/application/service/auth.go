package service

type AuthService interface {
	Generate(userID string) (string, error)
	GenerateRefresh(userID string) (string, error)
	Verify(token string) error
	GetSubject(token string) (string, error)
}
