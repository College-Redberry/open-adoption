package saveimages

import (
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/pet/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/pet/application/usecase/save_images"
)

type SaveImages struct {
	usecase command.UsecaseWithNoReturn[usecase.Input]
}

func New(usecase command.UsecaseWithNoReturn[usecase.Input]) *SaveImages {
	return &SaveImages{
		usecase: usecase,
	}
}

func (adopt *SaveImages) Handle(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return err
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	err = adopt.usecase.Execute(usecase.Input{
		ID:     id,
		Images: files,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
