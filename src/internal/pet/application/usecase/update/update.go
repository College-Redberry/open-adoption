package update

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
	err := pet.Gender(input.Gender).Validate()
	if err != nil {
		return err
	}

	err = usecase.petRepo.Update(input.ID, pet.PetProps{
		Name:        input.Name,
		Description: input.Description,
		Breed:       input.Breed,
		Age:         input.Age,
		Gender:      input.Gender,
	})
	if err != nil {
		return errs.ErrInternal.Wrap("failed to update pet", err)
	}

	return nil
}
