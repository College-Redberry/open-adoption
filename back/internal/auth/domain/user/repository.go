package user

type UserRepo interface {
	Create(user User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
