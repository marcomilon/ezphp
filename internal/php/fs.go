package php

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

// FindExec searchs for executable in local path environment
func FindExec(path string) (string, error) {

	path, err := exec.LookPath(path)
	if err != nil {
		return "", err
	}

	return path, nil

}

func FindPHP(srv Server) (string, error) {

	absPath, err := filepath.Abs(filepath.Dir(srv.Exec))
	if err != nil {
		return "", err
	}

	localPHP := absPath + string(os.PathSeparator) + path.Base(srv.Exec)

	if _, err = os.Stat(localPHP); os.IsNotExist(err) {
		return "", err
	}

	return localPHP, nil

}
