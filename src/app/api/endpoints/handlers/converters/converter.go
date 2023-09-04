package converters

import (
	"echofy_backend/src/core/errors"

	"github.com/google/uuid"
)

func ConvertFromStringToUUID(value, fieldName, fieldErrorDescription, errorMessage string) (*uuid.UUID, errors.Error) {
	convertedValue, err := uuid.Parse(value)
	if err != nil {
		return nil, getSingleValidarionErr(fieldName, fieldErrorDescription, errorMessage)
	}

	return &convertedValue, nil
}

func getSingleValidarionErr(fieldName, fieldErrorDescription, errorMessage string) *errors.ValidationError {
	invalidField := errors.InvalidField{Name: fieldName, Description: fieldErrorDescription}
	return errors.NewValidationError(errorMessage, invalidField)
}
