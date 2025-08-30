package login

import (
	"encoding/json"
	"net/http"

	command "github.com/college-redberry/open-adoption/internal/auth/application/usecase"
	usecase "github.com/college-redberry/open-adoption/internal/auth/application/usecase/login"
)

type Login struct {
	Usecase command.Usecase[usecase.Input, usecase.Output]
}

func New(usecase command.Usecase[usecase.Input, usecase.Output]) *Login {
	return &Login{
		Usecase: usecase,
	}
}

func (login *Login) Handle(w http.ResponseWriter, r *http.Request) error {
	var body Input
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	result, err := login.Usecase.Execute(usecase.Input{
		Email:    body.Email,
		Password: body.Password,
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
