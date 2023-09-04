package errors

var _ Error = NotFoundError{}

type NotFoundError struct {
	*superError
	message string
	err     error
}

func (instance NotFoundError) Error() string {
	return instance.err.Error()
}

func (instance NotFoundError) FriendlyMessage() string {
	return instance.message
}

func (instance NotFoundError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func (instance NotFoundError) LogLevel() int {
	return ErrorLevel
}

func NewNotFoundError(message string, err error) *NotFoundError {
	return &NotFoundError{
		superError: newSuperError(),
		message:    message,
		err:        err,
	}
}