package error

import "fmt"

// DomainError is domain error type
type DomainError struct {
	message string
	err     error
}

// NewDomainError returns new domain error
func NewDomainError(message string, err error) DomainError {
	return DomainError{
		message: message,
		err:     err,
	}
}

func (de DomainError) Error() string {
	if de.err != nil {
		return fmt.Sprintf("%s: %v", de.message, de.err)
	}

	return de.message
}

// GetMessage returns error message
func (de DomainError) GetMessage() string {
	return de.message
}

// GetError returns wrapped error
func (de DomainError) GetError() error {
	return de.err
}
