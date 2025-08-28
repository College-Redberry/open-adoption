package logout

import (
	"net/http"
)

type Logout struct{}

func New() *Logout {
	return &Logout{}
}

func (logout *Logout) Handle(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   -1,
	})

	w.WriteHeader(http.StatusOK)

	return nil
}
