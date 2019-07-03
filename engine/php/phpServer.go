package php

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

type PhpServer struct {
	PhpExe  string
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

func NewPhpServer(host string, docRoot string) PhpServer {
	return PhpServer{
		PHP_EXECUTABLE,
		host,
		docRoot,
	}
}

func (s PhpServer) Serve(ioCom IOCom) {
	logrus.Info("Starting web server using " + s.PhpExe + " -S " + s.Host + " -t " + s.DocRoot)

	out := outMsg{out: ioCom.Stdout}
	err := errMsg{err: ioCom.Stdout}

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)
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
