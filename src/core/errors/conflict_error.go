package errors

import "echofy_backend/src/core/messages"

var _ Error = ConflictError{}

type ConflictError struct {
	*superError
	conflictingFields []string
}

func (instance ConflictError) Error() string {
	return messages.ConflictErrorMessage
}

func (instance ConflictError) FriendlyMessage() string {
	return messages.ConflictErrorMessage
}

func (instance ConflictError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func (instance ConflictError) ConflictingFields() []string {
	return instance.conflictingFields
}

func NewConflictError(conflictingFields ...string) *ConflictError {
	return &ConflictError{
		superError:        newSuperError(),
		conflictingFields: conflictingFields,
	}
}
