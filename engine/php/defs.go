package php

type IOCom struct {
	Stdout chan string
	Stderr chan string
	Stdin  chan string
	Done   chan bool
}

type Installer interface {
	Install(iocom IOCom)
}

type Server interface {
	Serve(iocom IOCom)
	GetDocRoot() string
	GetHost() string
}
