package php

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cavaliercoder/grab"
	"github.com/marcomilon/ezphp/engine/ezio"
)

func (i Installer) Install(w ezio.EzIO) error {
	i.download()
	return nil
}

func (i Installer) WhereIs() (string, error) {
	var phpPath string
	var err error

	phpPath, err = whereIsGlobalPHP(PHP_EXECUTABLE)
	if err != nil {
		phpPath, err = whereIsLocalPHP(PHP_EXECUTABLE, i.Destination)
	}

	return phpPath, err
}

func (i Installer) download() (*grab.Response, error) {
	resp, err := grab.Get(i.Destination + i.Version, i.Source+i.Version)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func whereIsGlobalPHP(phpExe string) (string, error) {
	return exec.LookPath(phpExe)
}

func whereIsLocalPHP(phpExe string, target string) (string, error) {

	var err error

	if _, err = os.Stat(target); err == nil {
		absPath, _ := filepath.Abs(filepath.Dir(target))
		defaultExecPath := absPath + string(os.PathSeparator) + target + string(os.PathSeparator) + phpExe
		return defaultExecPath, nil
	}

	return "", err
}
