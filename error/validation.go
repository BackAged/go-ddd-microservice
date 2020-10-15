package error

import "encoding/json"

// ValidationError represents Validation Error
type ValidationError map[string]interface{}

// NewValidationError returns new validation error
func NewValidationError() ValidationError {
	return ValidationError{}
}

func (ve ValidationError) Error() string {
	buf, _ := json.Marshal(ve)
	return string(buf)
}

// Add adds a new error
func (ve ValidationError) Add(key, msg string) {
	ve[key] = msg
}

// HasErrors returns if it contains error
func (ve ValidationError) HasErrors() bool {
	return len(ve) != 0
}

// GetErrors returns errors
func (ve ValidationError) GetErrors() map[string]interface{} {
	return ve
}
