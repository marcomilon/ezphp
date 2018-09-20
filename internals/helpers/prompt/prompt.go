package prompt

import (
	"fmt"
	"strings"

	"github.com/marcomilon/ezphp/internals/helpers/output"
)

func Confirm(question string) bool {

	var confirmation string

	output.Info(fmt.Sprintf("%s [y/N]? ", question))
	fmt.Scan(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		return true
	}

	return false
}
