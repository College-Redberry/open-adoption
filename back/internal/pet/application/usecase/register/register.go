package register

import (
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
)

type Register struct {
	petRepo pet.PetRepo
}

func New(petRepo pet.PetRepo) *Register {
	return &Register{
		petRepo: petRepo,
	}
}

func (usecase *Register) Execute(input Input) (Output, error) {
	err := pet.Gender(input.Gender).Validate()
	if err != nil {
		return Output{}, err
	}

	newPet := pet.New(pet.PetProps{
		Name:   input.Name,
		Breed:  input.Breed,
		Age:    input.Age,
		Gender: input.Gender,
	})

	err = usecase.petRepo.Create(newPet)
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to register pet", err)
	}

	return Output{
		ID: newPet.ID,
	}, nil
}
