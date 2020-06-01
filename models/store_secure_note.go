package models

//SecureMessageRequest data model struct
type SecureMessageRequest struct {
	Note                 string `json:"message" validate:"required"`
	SelfDestruct         int    `json:"self_destruct" validate:"oneof=0 1 3 6 12 24 48 72 168 720"`
	DestructAfterOpening bool   `json:"destruct_after_opening" validate:"omitempty"`
}

//SecureMessageResponse data model struct
type SecureMessageResponse struct {
	Id  string `json:"id"`
	URL string `json:"url"`
}

//SecureMessage data model struct
type SecureMessage struct {
	Note                 []byte
	DestructAfterOpening bool
}
