package php

type IOCom struct {
	Stdout chan string
	Stderr chan string
	Stdin  chan string
	Done   chan bool
}

type Arguments struct {
	Host    string
	DocRoot string
}

type Installer interface {
	Install(iocom IOCom)
}

type Server interface {
	Serve(phpExe string, iocom IOCom)
}
