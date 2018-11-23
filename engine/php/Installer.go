package php

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Installer struct {
	Source      string
	Destination string
	Version     string
}

func (i Installer) Download(w io.Writer) error {
	fmt.Fprint(w, "Downloading...\n")

	return nil
}

func (i Installer) Install(w io.Writer) error {
	fmt.Fprint(w, "Installing...\n")

	return nil
}

func (i Installer) WhereIs(w io.Writer) (string, error) {
	var php string
	var err error

	fmt.Fprint(w, "Search for php...\n")
	php, err = whereIsGlobalPHP(PHP_EXECUTABLE)
	if err != nil {
		fmt.Fprint(w, "Search for php in %s...\n", i.Destination)
		php, err = whereIsLocalPHP(PHP_EXECUTABLE, i.Destination)
	}

	fmt.Fprint(w, fmt.Sprintf("Php found in...%s\n", php))
	return php, err
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
