package middleware

import (
	"net/http"
	"strings"

	"github.com/college-redberry/open-adoption/internal/auth/application/service"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

type Middleware struct {
	authService service.AuthService
}

func New(authService service.AuthService) *Middleware {
	return &Middleware{
		authService: authService,
	}
}

func (m *Middleware) Handle() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwt := r.Header.Get("Authorization")
			splitToken := strings.TrimPrefix(jwt, "Bearer ")
			if splitToken == "" {
				http.Error(w, errs.ErrNotAuthorized.Error(), http.StatusUnauthorized)
				return
			}

			err := m.authService.Verify(splitToken)
			if err != nil {
				http.Error(w, errs.ErrForbidden.Error(), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
