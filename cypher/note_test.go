package cypher

import (
	"testing"
)

func TestEncryptMessage(t *testing.T) {
	note := []byte("It does not do well to dwell on dreams and forget to live")
	key := []byte("3Go2iV@Rb$y*OdIMP3anI89@0i%o!e%w")

	_, err := EncryptNote(note, key)
	if err != nil {
		t.Fatal(err)
	}

	invalidKey := []byte("I'm an invalid key, ha ha ...")
	_, err = EncryptNote(note, invalidKey)
	if err == nil {
		t.Fail()
	}
}

func TestDecryptMessage(t *testing.T) {
	note := []byte("But you know, happiness can be found even in the darkest of times, " +
		"if one only remembers to turn on the light.")
	key := []byte("3Go2iV@Rb$y*OdIMP3anI89@0i%o!e%w")

	encryptedNote, err := EncryptNote(note, key)
	if err != nil {
		t.Fatal(err)
	}

	decryptedNote, err := DecryptNote(encryptedNote, key)
	if err != nil {
		t.Fatal(err)
	}
	if string(decryptedNote) != string(note) {
		t.Fail()
	}

	invalidKey := []byte("It matters not what someone is born, but what they grow to be.")
	_, err = DecryptNote(encryptedNote, invalidKey)
	if err == nil {
		t.Fail()
	}

	wrongKey := []byte("mCz&g8q2EoFuq&t5RsB1!q1BPE9dFNW6")
	_, err = DecryptNote(encryptedNote, wrongKey)
	if err == nil {
		t.Fail()
	}
}
