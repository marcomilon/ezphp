package php

const STDOUT = "STDOUT"
const STDERR = "STDERR"
const STDIN = "STDIN"
const STDINSTALL = "STDINSTALL"

type IOMessage struct {
	IOContext string
	Msg       string
}

type IOCom struct {
	Outmsg  chan IOMessage
	Confirm chan string
	Done    chan bool
}

func NewStdout(msg string) IOMessage {
	return IOMessage{
		STDOUT,
		msg,
	}
}

func NewStderr(msg string) IOMessage {
	return IOMessage{
		STDERR,
		msg,
	}
}

func NewStdin(msg string) IOMessage {
	return IOMessage{
		STDIN,
		msg,
	}
}

func NewStdInstall(msg string) IOMessage {
	return IOMessage{
		STDINSTALL,
		msg,
	}
}
