package php

type IOMessage struct {
	IOContext string
	Msg       string
}

type IOCom struct {
	Outmsg  chan IOMessage
	Confirm chan string
	Done    chan bool
}

func NewIOMessage(ioContext string, msg string) IOMessage {
	return IOMessage{
		ioContext,
		msg,
	}
}
