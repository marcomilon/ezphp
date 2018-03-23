package main

import (
	"fmt"
	"github.com/marcomilon/ezphp/installer"
	"os/exec"
)

func main() {

	fmt.Printf("[EzPhp] Launching to EzPHP\n")
    fmt.Printf("[EzPhp] https://github.com/marcomilon/ezphp\n")

	path, err := searchPhpBin()
	if err != nil {
		fmt.Printf("[EzPhp] php not found\n")
		path, err = installer.Install()
	}

	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
        return
	}

    command := exec.Command(path, "-S", "localhost:8888")
    execErr := command.Run()
	if execErr != nil {
        fmt.Printf("[Error] Unable to execute PHP. %s\n", execErr.Error())
        fmt.Printf("[Error] php path is. %s\n", path)
	}
}

func searchPhpBin() (string, error) {
	path, err := exec.LookPath(installer.PhpExecutable)
	if err != nil {
		return "", err
	}

	return path, nil
}