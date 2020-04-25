package keyManager

import (
	"path/filepath"
	"strconv"
	"time"
)

func GetKeyId() string {
	t := time.Now()
	// todo: this seems stupid, find a better way...
	return strconv.Itoa(t.Year()) + strconv.Itoa(int(t.Month()))
}

func getKeyPath() (string, error) {
	p, err := filepath.Abs(StorageDirectory)
	if err != nil {
		return "", err
	}

	return p, nil
}
