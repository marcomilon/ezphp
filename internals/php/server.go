package php

import (
	"os"
	"os/exec"
)

func Serve(php string, host string, docRoot string) error {
	cmd := exec.Command(php, "-S", host, "-t", docRoot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
