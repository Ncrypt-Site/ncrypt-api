package models

type SecureMessageRequest struct {
	Message      string `json:"message" validate:"required"`
	SelfDestruct int    `json:"self_destruct" validate:"required,oneof=0 1 3 6 12 24 48 72 168 720"`
	Password     string `json:"password" validate:"omitempty,min=8,max=24"`
}

type SecureMessage struct {
	Message  []byte
	KeyId    string
	Password string
}
