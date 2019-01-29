package cli

import (
	"fmt"
	"strings"
)

type WhiteIO struct{}

func (WhiteIO) Write(b []byte) (int, error) {
	s := string(b[0:])
	fmt.Print(s)

	return len(b), nil
}

func (WhiteIO) Info(s string) {
	fmt.Print(s)
}

func (WhiteIO) Error(s string) {
	fmt.Print(s)
}

func (io WhiteIO) Confirm(question string) bool {

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
