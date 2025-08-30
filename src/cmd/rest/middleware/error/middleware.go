package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/college-redberry/open-adoption/cmd/rest/middleware"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

type Middleware struct{}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Handle(next middleware.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		status := http.StatusInternalServerError
		message := errs.ErrInternal.Error()

		var domainErr *errs.DomainError
		if errors.As(err, &domainErr) {
			switch domainErr {
			case errs.ErrNotAuthorized:
				status = http.StatusUnauthorized
				message = domainErr.Error()
			case errs.ErrForbidden:
				status = http.StatusForbidden
				message = domainErr.Error()
			case errs.ErrNotFound:
				status = http.StatusNotFound
				message = domainErr.Error()
			case errs.ErrInvalidData:
				status = http.StatusBadRequest
				message = domainErr.Error()
			default:
				message = domainErr.Error()
			}
		}

		fmt.Println(err.Error())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(middleware.ErrorResponse{Error: message})
	}
}
