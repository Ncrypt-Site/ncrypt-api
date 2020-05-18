package helpers

import (
	"github.com/go-playground/validator/v10"
	"ncrypt-api/models"
	"testing"
)

func TestFormatValidationErrorMessage(t *testing.T) {
	data := models.SecureMessageRequest{
		Note:                 "",
		SelfDestruct:         0,
		DestructAfterOpening: false,
	}
	v := validator.New()
	err := v.Struct(data)

	validationError := FormatValidationErrorMessage(err)

	if len(validationError) != 1 ||
		validationError[0] != "validation failed for field: Note. reason: required." {
		t.Fatalf("expected 1 error but got %v, more details are: %v", len(validationError), validationError)
	}
}
