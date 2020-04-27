package processors

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"ncrypt-api/cypher"
	"ncrypt-api/models"
	"time"
)

func StoreMessage(client *redis.Client, k models.Key, m models.SecureMessageRequest) (uuid.UUID, error) {
	encryptedMessage, err := cypher.EncryptMessage([]byte(m.Message), k.PublicKey)
	if err != nil {
		fmt.Println(err)
	}

	message := models.SecureMessage{
		Message: encryptedMessage,
		KeyId:   k.Id,
	}

	if len(m.Password) > 0 {
		passwordHash, err := cypher.HashPassword([]byte(m.Password))
		if err != nil {
			return uuid.UUID{}, err
		}

		message.Password = passwordHash
	}

	messageUuid, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
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
