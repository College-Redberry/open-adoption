package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/college-redberry/open-adoption/internal/auth/domain/user"
	querier "github.com/college-redberry/open-adoption/internal/auth/infra/persist/user"
	"github.com/google/uuid"
)

type UserRepo struct {
	querier querier.Querier
}

func New(db querier.DBTX) *UserRepo {
	return &UserRepo{
		querier: querier.New(db),
	}
}

func (repo *UserRepo) Create(user user.User) error {
	err := repo.querier.CreateUser(context.Background(), querier.CreateUserParams{
		ID:        uuid.MustParse(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     string(user.Email),
		Password:  string(user.Password),
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) GetByEmail(email string) (user.User, error) {
	result, err := repo.querier.GetUserByEmail(context.Background(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, nil
		}

		return user.User{}, err
	}

	return user.User{
		ID: result.ID.String(),
		UserProps: user.UserProps{
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     user.Email(result.Email),
			Password:  user.Password(result.Password),
		},
	}, nil
}

func (repo *UserRepo) GetByID(id string) (user.User, error) {
	result, err := repo.querier.GetUserByID(context.Background(), uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, nil
		}

		return user.User{}, err
	}

	return user.User{
		ID: result.ID.String(),
		UserProps: user.UserProps{
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     user.Email(result.Email),
			Password:  user.Password(result.Password),
		},
	}, nil
}
