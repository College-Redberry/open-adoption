package refresh

import (
	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
	"github.com/college-redberry/open-adoption/internal/auth/domain/user"
)

type Refresh struct {
	UserRepo    user.UserRepo
	AuthService service.AuthService
}

func New(userRepo user.UserRepo, authService service.AuthService) *Refresh {
	return &Refresh{
		UserRepo:    userRepo,
		AuthService: authService,
	}
}

func (usecase *Refresh) Execute(input Input) (Output, error) {
	userID, err := usecase.AuthService.GetSubject(input.Token)
	if err != nil {
		return Output{}, errs.ErrInvalidData.Wrap("error to get subject", err)
	}

	user, err := usecase.UserRepo.GetByID(userID)
	if err != nil {
		return Output{}, errs.ErrInvalidData.Wrap("error to get user info", err)
	}

	if user.ID == "" {
		return Output{}, errs.ErrInvalidData.Wrap("error to get user info", nil)
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
