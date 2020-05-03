package models

import "github.com/google/uuid"

type RetrieveNoteRequest struct {
	Id       uuid.UUID `json:"id" validate:"required"`
	Password string    `json:"password" validate:"omitempty,min=8,max=24"`
}

type RetrieveNoteResponse struct {
	Note string `json:"note"`
}
