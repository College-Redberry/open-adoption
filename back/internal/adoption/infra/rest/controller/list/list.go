package list

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/adoption/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/list"
)

type List struct {
	usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *List {
	return &List{
		usecase: usecase,
	}
}

func (list *List) Handle(w http.ResponseWriter, r *http.Request) error {
	result, err := list.usecase.Execute(nil)
	if err != nil {
		return err
	}

	output := make([]Request, len(result))
	for _, request := range result {
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

	return nil
}
