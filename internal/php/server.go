package php

import (
	"os/exec"
)

// Server holds the information for the php server
type Server struct {
	Exec    string
	Host    string
	DocRoot string
}

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

// Serve start a php server
func (s Server) Serve(ioCom IOCom) {

	out := outMsg{out: ioCom.Stdout}
	err := errMsg{err: ioCom.Stdout}

	cmd := exec.Command(s.Exec, "-S", s.Host, "-t", s.DocRoot)
	cmd.Stdout = out
	cmd.Stderr = err

	errCmd := cmd.Run()

	if errCmd != nil {
		ioCom.Stderr <- "Error: " + errCmd.Error()
		ioCom.Done <- true
	}

	ioCom.Done <- true

}
