package request

import (
	request "github.com/college-redberry/open-adoption/internal/adoption/domain/request"
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
	"github.com/college-redberry/open-adoption/internal/pet/domain/pet"
)

type Aprove struct {
	requestRepo request.Repo
	petRepo     pet.PetRepo
}

func New(requestRepo request.Repo, petService pet.PetRepo) *Aprove {
	return &Aprove{
		requestRepo: requestRepo,
		petRepo:     petService,
	}
}

func (usecase *Aprove) Execute(input Input) error {
	err := usecase.requestRepo.Aprove(input.ID)
	if err != nil {
		return errs.ErrInternal.Wrap("failed to to aprove request", err)
	}

	request, err := usecase.requestRepo.GetByID(input.ID)
	if err != nil {
		return errs.ErrInternal.Wrap("failed to to get request", err)
	}

	err = usecase.petRepo.AdoptById(request.PetID)
	if err != nil {
		return errs.ErrInternal.Wrap("failed to to adopt pet", err)
	}

	return nil
}
