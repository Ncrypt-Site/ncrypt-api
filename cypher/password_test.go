package cypher

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	_, err := HashPassword([]byte("Expelliarmus"))
	if err != nil {
		t.Fail()
	}
}

func TestValidatePassword(t *testing.T) {
	plainPassword := []byte("Expelliarmus")
	hashedPassword, _ := HashPassword(plainPassword)

	if r := ValidatePassword(plainPassword, []byte(hashedPassword)); !r {
		t.Fail()
	}

	if r := ValidatePassword([]byte("Incendio"), []byte(hashedPassword)); r {
		t.Fail()
	}
}
