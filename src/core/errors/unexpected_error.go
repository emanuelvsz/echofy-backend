package errors

var _ Error = UnexpectedError{}

type UnexpectedError struct {
	*superError
	message string
	err     error
}

func (instance UnexpectedError) Error() string {
	return instance.err.Error()
}

func (instance UnexpectedError) FriendlyMessage() string {
	return instance.message
}

func (instance UnexpectedError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func (instance UnexpectedError) LogLevel() int {
	return ErrorLevel
}

func NewUnexpectedError(message string, err error) *UnexpectedError {
	return &UnexpectedError{
		superError: newSuperError(),
		message:    message,
		err:        err,
	}
}
