package php

import (
	"fmt"
	"os"
	"os/exec"
)

// FindExec searchs for executable in local path environment
func FindExec(path string) (string, error) {

	path, err := exec.LookPath(path)
	if err != nil {
		return "", err
	}

	return path, nil

}

func FindPHPExec(installFolder string) (string, error) {
	var phpExe string
	var err error

	phpExe, err = FindExec(PHP_EXECUTABLE)
	if err != nil {
		phpExe, err = FindLocalPHP(installFolder)
		if err != nil {
			return "", err
		}
	}

	return phpExe, nil
}

func FindLocalPHP(installFolder string) (string, error) {

	phpExe := fmt.Sprintf("%v/%v", installFolder, PHP_EXECUTABLE)

	if _, err := os.Stat(phpExe); os.IsNotExist(err) {
		return "", err
	}

	return phpExe, nil
}
