package php

import "github.com/marcomilon/ezphp/engine/ezio"

type Installer struct {
	Source      string
	Destination string
	Version     string
}

type Server struct {
	PhpPath      string
	DocumentRoot string
	Port         int
}

type EzInstaller interface {
	Install(w ezio.EzIO) error
	WhereIs() (string, error)
}

type EzServe interface {
	Serve(w ezio.EzIO) error
}
