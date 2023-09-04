package errors

var _ Error = UnavailableServiceError{}

type UnavailableServiceError struct {
	*superError
	message string
	err     error
}

func (instance UnavailableServiceError) Error() string {
	return instance.err.Error()
}

func (instance UnavailableServiceError) FriendlyMessage() string {
	return instance.message
}

func (instance UnavailableServiceError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func NewUnavailableServiceError(message string, err error) *UnavailableServiceError {
	return &UnavailableServiceError{
		superError: newSuperError(),
		message:    message,
		err:        err,
	}
}
