package aprove

import (
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/adoption/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/adoption/application/usecase/aprove"
)

type Aprove struct {
	usecase command.UsecaseWithNoReturn[usecase.Input]
}

func New(usecase command.UsecaseWithNoReturn[usecase.Input]) *Aprove {
	return &Aprove{
		usecase: usecase,
	}
}

func (controller *Aprove) Handle(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	err := controller.usecase.Execute(usecase.Input{
		ID: id,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
