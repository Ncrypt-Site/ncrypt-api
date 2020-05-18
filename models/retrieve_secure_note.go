package models

import "github.com/google/uuid"

type RetrieveNoteRequest struct {
	Id uuid.UUID `json:"id" validate:"required"`
}

type RetrieveNoteResponse struct {
	Note string `json:"note"`
}
