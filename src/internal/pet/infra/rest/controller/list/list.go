package list

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/list"
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

	output := make([]Pet, len(result))
	for _, pet := range result {
		output = append(output, Pet{
			ID:       pet.ID,
			Name:     pet.Name,
			Breed:    pet.Breed,
			Age:      pet.Age,
			Gender:   pet.Gender,
			IsAdoped: pet.IsAdoped,
			Images:   pet.Images,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

	return nil
}
