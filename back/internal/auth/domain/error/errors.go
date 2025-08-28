package errs

var (
	ErrInvalidOperation = New("invalid operation")
	ErrNotAuthorized    = New("not authorized")
	ErrForbidden        = New("forbidden")
	ErrInvalidData      = New("invalid data")
	ErrInternal         = New("internal error")
	ErrNotFound         = New("not found")
)
