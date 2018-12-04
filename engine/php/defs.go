package php

import "github.com/marcomilon/ezphp/engine/ezio"

const (
	EzPHPVersion = "1.1.0"
	EzPHPWebsite = "https://github.com/marcomilon/ezphp"
	PHPVersion   = "7.0.0"
)

type Installer struct {
	DownloadUrl string
	Filename    string
	InstallDir  string
}

type Server struct {
	PhpExe  string
	Host    string
	DocRoot string
}

type Arguments struct {
	Host       string
	DocRoot    string
	InstallDir string
}

type EzInstaller interface {
	Install(w ezio.EzIO) error
}

type EzServe interface {
	Serve(w ezio.EzIO) error
}
