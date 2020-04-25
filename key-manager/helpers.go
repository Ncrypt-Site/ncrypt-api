package keyManager

import (
	"strconv"
	"time"
)

func GetKeyId() string {
	t := time.Now()
	// todo: this seems stupid, find a better way...
	return strconv.Itoa(t.Year()) + strconv.Itoa(int(t.Month()))
}
