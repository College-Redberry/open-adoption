package retrieve

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/auth/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/retrieve"
)

type Retrieve struct {
	usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *Retrieve {
	return &Retrieve{
		usecase: usecase,
	}
}

func (getUserByID *Retrieve) Handle(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	result, err := getUserByID.usecase.Execute(usecase.Input{
		ID: id,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Output{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
	})

	return nil
}
