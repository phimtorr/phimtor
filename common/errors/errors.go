package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

type SlugError struct {
	slug      string
	err       string
	errorType ErrorType
}

func (e SlugError) Slug() string {
	return e.slug
}

func (e SlugError) Error() string {
	return e.err
}

func (e SlugError) ErrorType() ErrorType {
	return e.errorType
}

func NewUnknownError(slug, err string) SlugError {
	return SlugError{slug: slug, err: err, errorType: ErrorTypeUnknown}
}

func NewAuthorizationError(slug, err string) SlugError {
	return SlugError{slug: slug, err: err, errorType: ErrorTypeAuthorization}
}

func NewIncorrectInputError(slug, err string) SlugError {
	return SlugError{slug: slug, err: err, errorType: ErrorTypeIncorrectInput}
}
