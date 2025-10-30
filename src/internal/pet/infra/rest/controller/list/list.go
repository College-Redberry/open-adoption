package list

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/list"
	"github.com/ritwickdey/querydecoder"
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
	var queries Input
	err := querydecoder.New(r.URL.Query()).Decode(&queries)
	if err != nil {
		return err
	}

	result, err := list.usecase.Execute(usecase.Input{
		Name:      queries.Name,
		Breed:     queries.Breed,
		Age:       queries.Age,
		Gender:    queries.Gender,
		IsAdopted: queries.IsAdopted,
		Offset:    queries.Offset,
		Limit:     queries.Limit,
	})
	if err != nil {
		return err
	}

	output := Output{
		Count: result.Count,
		Data:  make([]Pet, len(result.Data)),
	}

	for i, pet := range result.Data {
		output.Data[i] = Pet{
			ID:          pet.ID,
			Name:        pet.Name,
			Description: pet.Description,
			Breed:       pet.Breed,
			Age:         pet.Age,
			Gender:      pet.Gender,
			IsAdopted:   pet.IsAdoped,
			Images:      pet.Images,
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

	return nil
}
