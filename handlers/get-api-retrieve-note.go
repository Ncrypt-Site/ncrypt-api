package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
	"ncrypt-api/models"
	"ncrypt-api/processors"
	"net/http"
)

func (di DI) GetSecureNoteV1(c echo.Context) error {
	id := c.Param("id")
	password := c.Request().Header.Get("X-NCRYPT-NOTE-PASSWORD")

	noteId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"input failure",
				nil,
				[]string{"invalid id provided"},
				nil,
			),
		)
	}

	payload := models.RetrieveNoteRequest{
		Id:       noteId,
		Password: password,
	}

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"validation failed",
				nil,
				helpers.FormatValidationErrorMessage(err),
				nil,
			),
		)
	}

	note, err := processors.RetrieveSecureNote(di.RedisClient, di.Key, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helpers.BuildResponse(
				http.StatusInternalServerError,
				"unable to retrieve note",
				nil,
				[]string{err.Error()},
				nil,
			),
		)
	}

	response := models.RetrieveNoteResponse{Note: string(note)}

	return c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			http.StatusOK,
			"Message retrieved.",
			&response,
			nil,
			nil,
		),
	)
}
