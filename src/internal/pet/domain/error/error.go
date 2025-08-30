package errs

import "fmt"

type DomainError struct {
	message string
	err     error
}

func New(message string) *DomainError {
	return &DomainError{
		message: message,
	}
}

func (d *DomainError) Error() string {
	return d.message
}

func (d *DomainError) Wrap(message string, err error) error {
	return &DomainError{
		message: fmt.Sprintf("%s: %s", d.message, message),
		err:     err,
	}
}

func (d *DomainError) Unwrap() error {
	return d.err
}
