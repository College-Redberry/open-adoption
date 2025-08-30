package pet

import (
	"slices"

	errs "github.com/college-redberry/open-adoption/internal/auth/domain/error"
)

type Gender string

var genders = []Gender{"male", "female"}

func (gender Gender) Validate() error {
	if !slices.Contains(genders, gender) {
		return errs.ErrInvalidData.Wrap("gender not valid", nil)
	}

	return nil
}
