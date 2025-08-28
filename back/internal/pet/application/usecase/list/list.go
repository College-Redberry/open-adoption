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

func (usecase *Update) Execute(input Input) (Output, error) {
	pets, err := usecase.petRepo.List()
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to list pets", err)
	}

	output := make([]Pet, len(pets))
	for _, pet := range pets {
		output = append(output, Pet{
			ID:       pet.ID,
			Name:     pet.Name,
			Breed:    pet.Breed,
			Age:      pet.Age,
			Gender:   pet.Gender,
			IsAdoped: pet.IsAdoped,
		})
	}

	return output, nil
}
