package head

import (
	"fmt"
	"strings"
)

type Terminal struct{}

func (Terminal) Write(b []byte) (int, error) {
	s := string(b[0:])
	fmt.Print(s)

	return len(b), nil
}

func (Terminal) Info(s string) {
	fmt.Print(s)
}

func (Terminal) Error(s string) {
	fmt.Print(s)
}

func (io Terminal) Confirm(question string) bool {

	var confirmation string

	io.Info(fmt.Sprintf("%s [y/N]? ", question))
	fmt.Scanln(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		return true
	}

	return false
}
