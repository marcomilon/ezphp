package php

const (
	EzPHPVersion = "1.1.0"
	EzPHPWebsite = "https://github.com/marcomilon/ezphp"
	PHPVersion   = "7.0.0"
)

type Command interface {
	Execute()
}

type IOChannels struct {
	Outmsg chan string
	Errmsg chan string
	Done   chan bool
}

type Arguments struct {
	Host       string
	DocRoot    string
	InstallDir string
}
