package php

import (
	"os/exec"
)

type PhpServer struct {
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

func NewPhpServer(arguments Arguments) PhpServer {
	return PhpServer{
		arguments.Host,
		arguments.DocRoot,
	}
}

func (s PhpServer) Serve(phpExe string, ioCom IOCom) {

	out := outMsg{out: ioCom.Stdout}
	err := errMsg{err: ioCom.Stdout}

	cmd := exec.Command(phpExe, "-S", s.Host, "-t", s.DocRoot)
	cmd.Stdout = out
	cmd.Stderr = err

	errCmd := cmd.Run()

	if errCmd != nil {
		ioCom.Stderr <- errCmd.Error()
		ioCom.Done <- true
	}
}

func (s PhpServer) GetDocRoot() string {
	return s.DocRoot
}

func (s PhpServer) GetHost() string {
	return s.Host
}
