package errors

import (
	"reflect"
)

var _ Error = ValidationError{}

type ValidationError struct {
	*superError
	message       string
	invalidFields []InvalidField
}

type InvalidField struct {
	Name        string
	Description string
}

func (instance ValidationError) Error() string {
	return instance.message
}

func (instance ValidationError) FriendlyMessage() string {
	return instance.message
}

func (instance ValidationError) Equals(err Error) bool {
	errorToCompare, ok := err.(*ValidationError)
	return ok && instance.message == errorToCompare.message && reflect.DeepEqual(instance.invalidFields,
		errorToCompare.invalidFields)
}

func (instance ValidationError) InvalidFields() []InvalidField {
	return instance.invalidFields
}

func NewValidationError(message string, invalidFields ...InvalidField) *ValidationError {
	return &ValidationError{
		superError:    newSuperError(),
		message:       message,
		invalidFields: invalidFields,
	}
}