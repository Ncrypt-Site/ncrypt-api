package models

import (
	"github.com/google/uuid"
	"time"
)

// StorageInterface defines the needed methods and their signatures for storage implementation
type StorageInterface interface {
	BuildConfiguration(c Config) (StorageInterface, error)
	Store(id uuid.UUID, data []byte, duration time.Duration) error
	Exists(id uuid.UUID) bool
	Retrieve(id uuid.UUID) (SecureMessage, error)
	Delete(id uuid.UUID) error
}
