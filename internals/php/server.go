package php

import (
	"os/exec"

	"github.com/marcomilon/ezphp/internals/helpers/ezio"
)

func Serve(php string, host string, docRoot string) error {

	ezOut := ezio.EzOut{Prompt: " EzPHP"}

	cmd := exec.Command(php, "-S", host, "-t", docRoot)
	cmd.Stdout = ezOut
	cmd.Stderr = ezOut
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
