package create

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/adoption/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/create"
)

type Create struct {
	usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *Create {
	return &Create{
		usecase: usecase,
	}
}

func (create *Create) Handle(w http.ResponseWriter, r *http.Request) error {
	var body Input
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	result, err := create.usecase.Execute(usecase.Input{
		PetID:                             body.PetID,
		Name:                              body.Name,
		Email:                             body.Email,
		Phone:                             body.Phone,
		Age:                               body.Age,
		HouseHoldAgreed:                   body.HouseHoldAgreed,
		AlreadyPets:                       body.AlreadyPets,
		AlreadyPetsCastratedAndVaccinated: body.AlreadyPetsCastratedAndVaccinated,
		Property:                          body.Property,
		OwnProperty:                       body.OwnProperty,
		Address:                           body.Address,
		Income:                            body.Income,
		SuitableLocation:                  body.SuitableLocation,
		AccessToTheStreet:                 body.AccessToTheStreet,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Output{
		ID: result.ID,
	})

	return nil
}
