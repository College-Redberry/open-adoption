package middleware

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/college-redberry/open-adoption/cmd/rest/middleware"
	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

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
			logger.Error("%s: %s", err.Error(), domainErr.Unwrap().Error())

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
		} else {
			logger.Error(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(middleware.ErrorResponse{Error: message})
	}
}
