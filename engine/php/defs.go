package php

import "github.com/marcomilon/ezphp/engine/ezio"

const (
	EzPHPVersion = "1.1.0"
	EzPHPWebsite = "https://github.com/marcomilon/ezphp"
	PHPVersion   = "7.0.0"
)

type Channels struct {
	Outmsg chan string
	Errmsg chan string
	Done   chan bool
}

type Installer struct {
	DownloadUrl string
	Filename    string
	InstallDir  string
	Channels
}

type Server struct {
	PhpExe  string
	Host    string
	DocRoot string
	Channels
}

type Arguments struct {
	Host       string
	DocRoot    string
	InstallDir string
}

type EzInstaller interface {
	Install()
}

type EzServe interface {
	Serve(w ezio.EzIO) error
}
