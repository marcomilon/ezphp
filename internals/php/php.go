package php

import (
	"os/exec"

	"github.com/marcomilon/ezphp/internals/helpers/ezio"
)

func Cli(php string, arg string) error {
	cmd := exec.Command(php, arg)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Serve(php string, host string, docRoot string) error {
	ezOut := ezio.EzOut{Prompt: " Serve"}

	cmd := exec.Command(php, "-S", host, "-t", docRoot)
	cmd.Stdout = ezOut
	cmd.Stderr = ezOut
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
