package helpers

import (
	"io/ioutil"
	"os"
)

func WriteToFile(fileName string, b []byte, permission os.FileMode) error {
	err := ioutil.WriteFile(fileName, b, permission)
	if err != nil {
		return err
	}

	return nil
}
