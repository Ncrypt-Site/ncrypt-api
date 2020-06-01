package models

import "github.com/google/uuid"

//RetrieveNoteRequest data model struct
type RetrieveNoteRequest struct {
	Id uuid.UUID `json:"id" validate:"required"`
}

//RetrieveNoteResponse data model struct
type RetrieveNoteResponse struct {
	Note string `json:"note"`
}
