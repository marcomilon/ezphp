package docroot

import (
	"fmt"
	"os"
)

func Create(path string) error {
	return os.MkdirAll(path, 0755)
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIndex(path, template string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	fmt.Fprintf(file, template)

	return nil
}
