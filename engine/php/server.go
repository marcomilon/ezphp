package php

import (
	"io"
	"os/exec"
)

func (s Server) Serve(stdout io.Writer, stderr io.Writer) error {

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
