package models

import "github.com/google/uuid"

//RetrieveNoteRequest
type RetrieveNoteRequest struct {
	Id uuid.UUID `json:"id" validate:"required"`
}

//RetrieveNoteResponse
type RetrieveNoteResponse struct {
	Note string `json:"note"`
}
