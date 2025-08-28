package retrieve

import (
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
	"github.com/college-redberry/open-adoption/internal/auth/domain/user"
)

type Retrieve struct {
	userRepo user.UserRepo
}

func New(userRepo user.UserRepo) *Retrieve {
	return &Retrieve{
		userRepo: userRepo,
	}
}

func (usecase *Retrieve) Execute(input Input) (Output, error) {
	user, err := usecase.userRepo.GetByID(input.ID)
	if err != nil {
		return Output{}, errs.ErrNotFound.Wrap("user not found", err)
	}

	return Output{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     string(user.Email),
	}, nil
}
