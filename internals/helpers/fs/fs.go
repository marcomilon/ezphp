package fs

import (
	"os"
	"os/exec"
	"path/filepath"
)

func WhereIsGlobalPHP(phpExe string) (string, error) {
	return exec.LookPath(phpExe)
}

func WhereIsLocalPHP(phpExe string, target string) (string, error) {

	var err error

	if _, err = os.Stat(target); err == nil {
		absPath, _ := filepath.Abs(filepath.Dir(target))
		defaultExecPath := absPath + string(os.PathSeparator) + target + string(os.PathSeparator) + phpExe
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
