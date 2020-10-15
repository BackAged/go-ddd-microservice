package error

// NotFoundError service error
type NotFoundError struct {
	message string
	errors  []string
}

func (anf *NotFoundError) Error() string {
	return anf.message
}

// Add adds new error
func (anf *NotFoundError) Add(value string) {
	anf.errors = append(anf.errors, value)
}

// GetMessage returns error message
func (anf *NotFoundError) GetMessage() string {
	return anf.message
}

// GetErrors returns erros
func (anf *NotFoundError) GetErrors() []string {
	return anf.errors
}

// NewNotFoundError returns new NotFound error
func NewNotFoundError(message string, errors []string) *NotFoundError {
	return &NotFoundError{
		message: message,
		errors:  errors,
	}
}
