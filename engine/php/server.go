package php

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

type outMsg struct {
	out chan string
}

type errMsg struct {
	err chan string
}

func (o outMsg) Write(p []byte) (n int, err error) {
	s := string(p)
	o.out <- s

	return len(p), nil
}

func (e errMsg) Write(p []byte) (n int, err error) {
	s := string(p)
	e.err <- s

	return len(p), nil
}

func (s Server) Serve() {
	logrus.Info("Starting web server using " + s.PhpExe + " -S " + s.Host + " -t " + s.DocRoot)

	out := outMsg{out: s.Outmsg}
	err := errMsg{err: s.Errmsg}

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)
	cmd.Stdout = out
	cmd.Stderr = err

	errCmd := cmd.Run()
	
	if errCmd != nil {
		s.Errmsg <- errCmd.Error()
		s.Done <- true
	}
}
