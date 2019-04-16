package php

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

type Server struct {
	PhpExe  string
	Host    string
	DocRoot string
}

type outMsg struct {
	out chan IOMessage
}

type errMsg struct {
	err chan IOMessage
}

func (o outMsg) Write(p []byte) (n int, err error) {
	s := string(p)

	outmsg := NewIOMessage("stdout", s)
	o.out <- outmsg

	return len(p), nil
}

func (e errMsg) Write(p []byte) (n int, err error) {
	s := string(p)

	errmsg := NewIOMessage("stderr", s)
	e.err <- errmsg

	return len(p), nil
}

func (s Server) StartServer(ioCom IOCom) {
	logrus.Info("Starting web server using " + s.PhpExe + " -S " + s.Host + " -t " + s.DocRoot)

	out := outMsg{out: ioCom.Outmsg}
	err := errMsg{err: ioCom.Outmsg}

	cmd := exec.Command(s.PhpExe, "-S", s.Host, "-t", s.DocRoot)
	cmd.Stdout = out
	cmd.Stderr = err

	errCmd := cmd.Run()

	if errCmd != nil {
		errmsg := NewIOMessage("stderr", errCmd.Error())
		ioCom.Outmsg <- errmsg
		ioCom.Done <- true
	}
}
