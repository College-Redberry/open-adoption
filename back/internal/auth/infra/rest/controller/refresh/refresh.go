package refresh

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/auth/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/refresh"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

type Refresh struct {
	Usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *Refresh {
	return &Refresh{
		Usecase: usecase,
	}
}

func (refresh *Refresh) Handle(w http.ResponseWriter, r *http.Request) error {
	refreshToken, err := r.Cookie("refreshToken")
	if err != nil {
		return err
	}

	if refreshToken.Value == "" {
		return errs.ErrNotAuthorized
	}

	result, err := refresh.Usecase.Execute(usecase.Input{
		Token: refreshToken.Value,
	})
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    result.RefreshToken,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   3600,
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Output{
		Token: result.Token,
	})

	return nil
}
