package fs

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/php"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func WhereIsPHP(installDir string) (string, error) {
	var phpPath string
	var err error

	phpPath, err = whereIsGlobalPHP(php.PHP_EXECUTABLE)
	if err != nil {
		phpPath, err = whereIsLocalPHP(php.PHP_EXECUTABLE, installDir)
	}

	log.Println("PHP founded in " + phpPath)

	return phpPath, err
}

func whereIsGlobalPHP(phpExe string) (string, error) {
	log.Println("Searching for PHP in $PATH")
	return exec.LookPath(phpExe)
}

func whereIsLocalPHP(phpExe string, target string) (string, error) {
	var err error
	absPath, _ := filepath.Abs(filepath.Dir(target))
	localPHP := absPath + string(os.PathSeparator) + target + string(os.PathSeparator) + phpExe

	log.Println("Searching for PHP in " + localPHP)

	if _, err = os.Stat(localPHP); !os.IsNotExist(err) {
		return localPHP, nil
	}

	return "", err
}
