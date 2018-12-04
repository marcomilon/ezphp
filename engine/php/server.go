package php

import (
	"io"
	"log"
	"os/exec"
)

func (s Server) Serve(stdout io.Writer, stderr io.Writer) error {
	log.Println("Starting webserver using " + s.PhpExe + "-S" + s.Host + "-t" + s.DocRoot)

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

	if err != nil {
		log.Println("Failed to start webserver: " + err.Error())
		return err
	}

	return nil
}
