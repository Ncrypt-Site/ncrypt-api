package key_manager

import (
	"strconv"
	"time"
)

func getKeyId() string {
	t := time.Now()
	// todo: this seems stupid, find a better way...
	return strconv.Itoa(t.Year()) + strconv.Itoa(int(t.Month()))
}
