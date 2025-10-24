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
	pets, err := usecase.petRepo.List()
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to list pets", err)
	}

	output := make([]Pet, len(pets))
	for _, pet := range pets {
		images, err := usecase.getPetImageById(pet.ID)
		if err != nil {
			return Output{}, err
		}

		output = append(output, Pet{
			ID:       pet.ID,
			Name:     pet.Name,
			Breed:    pet.Breed,
			Age:      pet.Age,
			Gender:   pet.Gender,
			IsAdoped: pet.IsAdoped,
			Images:   images,
		})
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
