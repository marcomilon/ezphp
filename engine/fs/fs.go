package fs

import (
	"fmt"
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

		createDefaultIndex(dir)
	}
	return nil
}

func createDefaultIndex(basePath string) {
	log.Println("Creating default index.php in directory: " + basePath)

	file, err := os.Create(basePath + string(os.PathSeparator) + "index.php")
	if err != nil {
		log.Println("Cannot create default index.php:  " + err.Error())
		return
	}

	defer file.Close()

	fmt.Fprintf(file, "<?php\n")
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "echo \"Welcome to your personal web server.<br>Replace this file with your own index.php.<br>\";\n")
	fmt.Fprintf(file, "echo \"This file is located in directory: "+basePath+"\";\n")
	fmt.Fprintf(file, "\n")
}

func WhereIsPHP(installDir string) (string, error) {
	var phpPath string
	var err error

	phpPath, err = whereIsGlobalPHP(php.PHP_EXECUTABLE)
	if err != nil {
		phpPath, err = whereIsLocalPHP(php.PHP_EXECUTABLE, installDir)
	}

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
