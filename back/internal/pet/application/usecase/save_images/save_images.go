package request

import (
	"io"

	"github.com/college-redberry/open-adoption/internal/pet/application/service"
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
	"github.com/google/uuid"
)

type Aprove struct {
	storageService service.Storage
	petRepo        pet.PetRepo
}

func New(petRepo pet.PetRepo, storageService service.Storage) *Aprove {
	return &Aprove{
		petRepo:        petRepo,
		storageService: storageService,
	}
}

func (usecase *Aprove) Execute(input Input) error {
	fileNames := []string{}

	for index, image := range input.Images {
		fileName := uuid.Must(uuid.NewUUID())
		fileNames[index] = fileName.String()

		file, err := image.Open()
		if err != nil {
			return errs.ErrInternal.Wrap("error opening file", err)

		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return errs.ErrInternal.Wrap("error reading file", err)
		}

		err = usecase.storageService.UploadFile(fileName.String(), data)
		if err != nil {
			return errs.ErrInternal.Wrap("error to save image", err)
		}
	}

	err := usecase.petRepo.SaveImagesById(input.ID, fileNames)
	if err != nil {
		return errs.ErrInternal.Wrap("error saving images", err)
	}

	return nil
}
