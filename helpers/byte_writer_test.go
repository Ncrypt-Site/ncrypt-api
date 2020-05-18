package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	if os.Getuid() == 0 {
		t.Fatal("tests should not be run under root user")
	}

	data := []byte("Codes are, in my not so humble opinion, our most inexhaustible source of magic, " +
		"capable of both influencing injury, and remedying it.")

	absPath, _ := filepath.Abs(".")
	tempFile := absPath + "/tmp.txt"
	permission := os.FileMode(0644)

	err := WriteToFile(tempFile, data, permission)
	if err != nil {
		t.Fatalf("unable to write a file: %v", err)
	}

	state, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("unable to get file info: %v", err)
	}

	if state.Mode() != permission.Perm() {
		t.Fatalf("file should be with %v permission but rather stored with %v", state.Mode(), permission)
	}

	fileContent, err := ioutil.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("unable to read the file: %v", err)
	}

	if string(fileContent) != string(data) {
		t.Fatal("data inside the file doesn't match with the test data")
	}

	_ = os.Remove(tempFile)
}

func TestWriteToFileForFailure(t *testing.T) {
	if os.Getuid() == 0 {
		t.Fatal("tests should not be run under root user")
	}

	data := []byte("Codes are, in my not so humble opinion, our most inexhaustible source of magic, " +
		"capable of both influencing injury, and remedying it.")

	absPath, _ := filepath.Abs("../.invalid_dir")
	tempFile := absPath + "/tmp.txt"
	permission := os.FileMode(0000)

	err := WriteToFile(tempFile, data, permission)
	if err == nil {
		t.Fail()
	}

	_, err = os.Stat(tempFile)
	if err == nil {
		t.Fail()
	}
}
