package helpers

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrorMessage(err error) []string {
	var messages []string

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := "validation failed for field: " + err.Field() + ". reason: " + err.Tag() + "."
			if len(err.Param()) != 0 {
				msg += " additional data: " + err.Param()
			}

			messages = append(messages, msg)
		}
	}

	return messages
}
