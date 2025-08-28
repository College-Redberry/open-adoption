package update

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/update"
)

type Update struct {
	usecase command.UsecaseWithNoReturn[usecase.Input]
}

func New(usecase command.UsecaseWithNoReturn[usecase.Input]) *Update {
	return &Update{
		usecase: usecase,
	}
}

func (update *Update) Handle(w http.ResponseWriter, r *http.Request) error {
	var body Input
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	id := r.PathValue("id")

	err = update.usecase.Execute(usecase.Input{
		ID:     id,
		Name:   body.Name,
		Breed:  body.Breed,
		Age:    body.Age,
		Gender: body.Gender,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
