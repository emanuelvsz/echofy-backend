package errors

var _ Error = (*MissingInformationError)(nil)

type MissingInformationError struct {
	*superError
	message string
}

func (instance MissingInformationError) Error() string {
	return instance.Error()
}

func (instance MissingInformationError) FriendlyMessage() string {
	return instance.message
}

func (instance MissingInformationError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func NewMissingInformationError(message string) *MissingInformationError {
	return &MissingInformationError{
		superError: newSuperError(),
		message:    message,
	}
}