package php

import (
	"io"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func (s Server) Serve(stdout io.Writer, stderr io.Writer) error {
	logrus.Info("Starting web server using " + s.PhpExe + " -S " + s.Host + " -t " + s.DocRoot)

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

	if err != nil {
		logrus.Error("Failed to start web server: " + err.Error())
		return err
	}

	return nil
}
