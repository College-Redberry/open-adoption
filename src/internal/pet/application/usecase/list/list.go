package update

import (
	"github.com/college-redberry/open-adoption/internal/pet/application/service"
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
)

type Update struct {
	petRepo        pet.PetRepo
	storageService service.Storage
}

func New(petRepo pet.PetRepo, storageService service.Storage) *Update {
	return &Update{
		petRepo:        petRepo,
		storageService: storageService,
	}
}

func (usecase *Update) Execute(input Input) (Output, error) {
	count, err := usecase.petRepo.Count(pet.Filters{
		Name:      input.Name,
		Breed:     input.Breed,
		Age:       input.Age,
		Gender:    input.Gender,
		IsAdopted: input.IsAdopted,
	})
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to count pets", err)
	}

	if count == 0 {
		return Output{}, nil
	}

	pets, err := usecase.petRepo.List(pet.Filters{
		Name:      input.Name,
		Breed:     input.Breed,
		Age:       input.Age,
		Gender:    input.Gender,
		IsAdopted: input.IsAdopted,
		Offset:    input.Offset,
		Limit:     input.Limit,
	})
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to list pets", err)
	}

	output := Output{
		Count: count,
		Data:  make([]Pet, len(pets)),
	}

	for i, pet := range pets {
		images, err := usecase.getPetImageById(pet.ID)
		if err != nil {
			return Output{}, err
		}

		output.Data[i] = Pet{
			ID:          pet.ID,
			Name:        pet.Name,
			Description: pet.Description,
			Breed:       pet.Breed,
			Age:         pet.Age,
			Gender:      pet.Gender,
			IsAdoped:    pet.IsAdoped,
			Images:      images,
		}
	}

	return output, nil
}

func (usecase *Update) getPetImageById(id string) ([]string, error) {
	imagesIDs, err := usecase.petRepo.ListImagesById(id)
	if err != nil {
		return nil, errs.ErrInternal.Wrap("failed to get list images", err)
	}

	if len(imagesIDs) == 0 {
		return nil, nil
	}

	images := make([]string, len(imagesIDs))
	for i, imagesID := range imagesIDs {
		image, err := usecase.storageService.GenerateUrl(imagesID)
		if err != nil {
			return nil, errs.ErrInternal.Wrap("failed to get pet image", err)
		}

		images[i] = image
	}

	return images, nil
}
