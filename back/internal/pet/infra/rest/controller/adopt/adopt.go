package adopt

import (
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/adopt"
)

type Adopt struct {
	usecase command.UsecaseWithNoReturn[usecase.Input]
}

func New(usecase command.UsecaseWithNoReturn[usecase.Input]) *Adopt {
	return &Adopt{
		usecase: usecase,
	}
}

func (adopt *Adopt) Handle(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	err := adopt.usecase.Execute(usecase.Input{
		ID: id,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
