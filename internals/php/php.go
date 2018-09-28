package php

import (
	"io"
	"os/exec"
)

func Run(php string, arg []string, stdout io.Writer, stderr io.Writer) error {
	arguments := arg[1:]
	cmd := exec.Command(php, arguments...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
