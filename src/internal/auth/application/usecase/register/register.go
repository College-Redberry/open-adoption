package register

import (
	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
	"github.com/college-redberry/open-adoption/internal/auth/domain/user"
)

type Register struct {
	userRepo       user.UserRepo
	hashService    service.HashService
	encryptService service.EncryptService
}

func New(userRepo user.UserRepo, hashService service.HashService, encryptService service.EncryptService) *Register {
	return &Register{
		userRepo:       userRepo,
		hashService:    hashService,
		encryptService: encryptService,
	}
}

func (usecase *Register) Execute(input Input) (Output, error) {
	err := user.Email(input.Email).Validate()
	if err != nil {
		return Output{}, err
	}

	decryptedPassword, err := usecase.encryptService.Decrypt([]byte(input.Password))
	if err != nil {
		return Output{}, err
	}
	input.Password = string(decryptedPassword)

	err = user.Password(input.Password).Validate()
	if err != nil {
		return Output{}, err
	}

	existingUser, err := usecase.userRepo.GetByEmail(input.Email)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to retrieve user", err)
	}

	if existingUser.ID != "" {
		return Output{}, errs.ErrInvalidData.Wrap("already used email", err)
	}

	hashedPassword, err := usecase.hashService.Hash(input.Password)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to validate password", err)
	}

	newUser := user.New(user.UserProps{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     user.Email(input.Email),
		Password:  user.Password(hashedPassword),
	})

	err = usecase.userRepo.Create(newUser)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to create user", err)
	}

	return Output{ID: newUser.ID}, nil
}
