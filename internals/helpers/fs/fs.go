package fs

import (
	"os"
	"os/exec"
	"path/filepath"
)

const local = "phpbin"

func WhereIsGlobalPHP(phpExe string) (string, error) {
	return exec.LookPath(phpExe)
}

func WhereIsLocalPHP(phpExe string) (string, error) {

	var err error

	if _, err = os.Stat(local); err == nil {
		absPath, _ := filepath.Abs(filepath.Dir(local))
		defaultExecPath := absPath + string(os.PathSeparator) + local + string(os.PathSeparator) + phpExe
		return defaultExecPath, nil
	}

	return "", err
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
