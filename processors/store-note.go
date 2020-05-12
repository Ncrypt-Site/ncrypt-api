package processors

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/models"
	"time"
)

func StoreMessage(client *redis.Client, m models.SecureMessageRequest) (uuid.UUID, error) {
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

	err = client.Set(messageUuid.String(), jsonData, time.Hour*time.Duration(m.SelfDestruct)).Err()
	if err != nil {
		return uuid.UUID{}, err
	}

	return messageUuid, nil
}
