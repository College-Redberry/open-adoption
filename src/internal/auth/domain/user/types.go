package user

import (
	"regexp"

	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

type Email string

func (email Email) Validate() error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(string(email)) {
		return errs.ErrInvalidData.Wrap("invalid email", nil)
	}
	return nil
}

type Password string

func (password Password) Validate() error {
	passStr := string(password)

	if len(passStr) < 8 {
		return errs.ErrInvalidData.Wrap("password must be at least 8 characters long", nil)
	}

	specialCharRegex := regexp.MustCompile(`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	if !specialCharRegex.MatchString(passStr) {
		return errs.ErrInvalidData.Wrap("password must contain at least one special character", nil)
	}

	numberRegex := regexp.MustCompile(`[0-9]`)
	if !numberRegex.MatchString(passStr) {
		return errs.ErrInvalidData.Wrap("password must contain at least one number", nil)
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(passStr) {
		return errs.ErrInvalidData.Wrap("password must contain at least one uppercase letter", nil)
	}

	return nil
}
