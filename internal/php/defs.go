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

type Install interface {
	Install(iocom IOCom) error
}

type Serve interface {
	Serve(iocom IOCom)
}
