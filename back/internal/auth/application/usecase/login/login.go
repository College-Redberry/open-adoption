package login

import (
	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
	"github.com/college-redberry/open-adoption/internal/auth/domain/user"
)

type Login struct {
	UserRepo    user.UserRepo
	AuthService service.AuthService
	HashService service.HashService
}

func New(userRepo user.UserRepo, authService service.AuthService, hashService service.HashService) *Login {
	return &Login{
		UserRepo:    userRepo,
		AuthService: authService,
		HashService: hashService,
	}
}

func (usecase *Login) Execute(input Input) (Output, error) {
	user, err := usecase.UserRepo.GetByEmail(input.Email)
	if err != nil {
		return Output{}, errs.ErrInvalidData.Wrap("error to get user info", err)
	}

	if user.ID == "" {
		return Output{}, errs.ErrInvalidData.Wrap("error to get user info", nil)
	}

	isPasswordCorrect, err := usecase.HashService.Verify(input.Password, string(user.Password))
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("error to validate password", err)
	}

	if !isPasswordCorrect {
		return Output{}, errs.ErrInvalidData.Wrap("email and password mismatch", nil)
	}

	token, err := usecase.AuthService.Generate(user.ID)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to generate token", err)
	}

	refreshToken, err := usecase.AuthService.GenerateRefresh(user.ID)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to generate refresh token", err)
	}

	return Output{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
