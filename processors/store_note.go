package processors

import (
	"encoding/json"
	"github.com/google/uuid"
	"ncrypt-api/models"
	"time"
)

//StoreMessage handles storing a secure note
func StoreMessage(storage models.StorageInterface, m models.SecureMessageRequest) (uuid.UUID, error) {
	message := models.SecureMessage{
		Note:                 []byte(m.Note),
		DestructAfterOpening: m.DestructAfterOpening,
	}

	messageUuid, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}

	if m.SelfDestruct == 0 {
		m.SelfDestruct = 720
		message.DestructAfterOpening = true
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = storage.Store(messageUuid, jsonData, time.Hour*time.Duration(m.SelfDestruct))
	if err != nil {
		return uuid.UUID{}, err
	}

	return messageUuid, nil
}
