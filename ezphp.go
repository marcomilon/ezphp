package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/marcomilon/ezphp/installer"
	"github.com/marcomilon/ezphp/server"
	"os/exec"
	"runtime"
	"strings"
)

const windows = "linux"

func main() {

	var phpExecPath string
	var err error

	if runtime.GOOS == windows {
		phpExecPath = "php.exe"
	} else {
		phpExecPath = "/usr/bin/php232"
	}

	php := flag.String("php", phpExecPath, "Path to php executable")
	host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
	public := flag.String("public", "web", "Path to public directory")

	flag.Parse()

	phpExecPath, err = exec.LookPath(phpExecPath)
	if err != nil {
		fmt.Printf("[Error] php executable %s not found in path\n", phpExecPath)
		if runtime.GOOS == windows {
			phpExecPath, err = askToInstallPhp()
			if err != nil {
				fmt.Println("[Info] php won't be installed. bye bye")
				return
			}

            phpExecPath, err = installer.Install()
            if err != nil {
                fmt.Println("[Error] " + err.Error())
                return
            }
		}
	}

	args := server.Args{
		Php:    *php,
		Host:   *host,
		Public: *public,
	}

	server.Run(args)

	return
}

func askToInstallPhp() (string, error) {
	var confirmation string

	fmt.Print("[Installer] Would you like to install php locally [y/N]? ")
	fmt.Scan(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" || confirmation == "yes" {
		return "path", nil
	}

	return "", errors.New("Unable to cofirm php installation.")
}
