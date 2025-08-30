package register

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/register"
)

type Register struct {
	usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *Register {
	return &Register{
		usecase: usecase,
	}
}

func (register *Register) Handle(w http.ResponseWriter, r *http.Request) error {
	var body Input
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	result, err := register.usecase.Execute(usecase.Input{
		Name:   body.Name,
		Breed:  body.Breed,
		Age:    body.Age,
		Gender: body.Gender,
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
