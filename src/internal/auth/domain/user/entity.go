package user

import "github.com/google/uuid"

type UserProps struct {
	FirstName string
	LastName  string
	Email     Email
	Password  Password
}

type User struct {
	ID string
	UserProps
}

func New(user UserProps) User {
	return User{
		ID:        uuid.New().String(),
		UserProps: user,
	}
}
