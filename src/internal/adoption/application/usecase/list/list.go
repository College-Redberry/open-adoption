package list

import (
	adoption "github.com/college-redberry/open-adoption/internal/adoption/domain/request"
	errs "github.com/college-redberry/open-adoption/internal/pet/domain/error"
)

type Aprove struct {
	requestRepo adoption.Repo
}

func New(requestRepo adoption.Repo) *Aprove {
	return &Aprove{
		requestRepo: requestRepo,
	}
}

func (usecase *Aprove) Execute(input Input) (Output, error) {
	requests, err := usecase.requestRepo.List()
	if err != nil {
		return Output{}, errs.ErrInternal.Wrap("failed to list requests", err)
	}

	output := make([]Request, len(requests))
	for _, request := range requests {
		output = append(output, Request{
			ID:                                request.ID,
			PetID:                             request.PetID,
			Name:                              request.Name,
			Email:                             request.Email,
			Phone:                             request.Phone,
			ApprovedAt:                        request.ApprovedAt,
			Age:                               request.Age,
			HouseHoldAgreed:                   request.HouseHoldAgreed,
			AlreadyPets:                       request.AlreadyPets,
			AlreadyPetsCastratedAndVaccinated: request.AlreadyPetsCastratedAndVaccinated,
			Property:                          request.Property,
			OwnProperty:                       request.OwnProperty,
			Address:                           request.Address,
			Income:                            request.Income,
			SuitableLocation:                  request.SuitableLocation,
			AccessToTheStreet:                 request.AccessToTheStreet,
			CreatedAt:                         request.CreatedAt,
		})
	}

	return output, nil
}
