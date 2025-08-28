package adopt

import (
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
)

type Update struct {
	petRepo pet.PetRepo
}

func New(petRepo pet.PetRepo) *Update {
	return &Update{
		petRepo: petRepo,
	}
}

func (usecase *Update) Execute(input Input) error {
	err := usecase.petRepo.AdoptById(input.ID)
	if err != nil {
		return errs.ErrInternal.Wrap("failed to adopt pet", err)
	}

	return nil
}
