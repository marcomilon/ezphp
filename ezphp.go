package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/marcomilon/ezphp/installer"
	"github.com/marcomilon/ezphp/server"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const localPhpInstallDir = installer.PhpDir

func main() {

	var defaultExecPath = installer.PhpExecutable
	var err error

	banner()

	php := flag.String("php", "", "Path to php executable")
	host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
	public := flag.String("public", "public", "Path to public directory")

	flag.Parse()

	if *php == "" {

		defaultExecPath, err = searchForPhp(defaultExecPath, err)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("[Info] Press Enter to exit...")
			fmt.Scanln()
			fmt.Scanln()
			return
		}

		php = &defaultExecPath
	}

	fmt.Println("[Info] Copy your php documents on " + *public + " directory")
	installer.CreateDirIfNotExist(*public)

	args := server.Args{
		Php:    *php,
		Host:   *host,
		Public: *public,
	}

	server.Run(args)

	return
}

func searchForPhp(defaultExecPath string, err error) (string, error) {

	if _, err := os.Stat(localPhpInstallDir + string(os.PathSeparator) + defaultExecPath); err == nil {
		fmt.Println("[Info] Local php installation founded")
		absPath, _ := filepath.Abs(filepath.Dir(localPhpInstallDir))
		defaultExecPath = absPath + string(os.PathSeparator) + localPhpInstallDir + string(os.PathSeparator) + defaultExecPath
		return defaultExecPath, nil
	}

	defaultExecPath, err = exec.LookPath(defaultExecPath)
	if err != nil {
		fmt.Println("[Error] php executable not found in path")

		defaultExecPath, err = askToInstallPhp()
		if err != nil {
			return "", errors.New("[Info] php won't be installed. bye bye.")
		}

		defaultExecPath, err = installer.Install()
		if err != nil {
			return "", errors.New("[Error] " + err.Error())
		}
	}

	return defaultExecPath, nil
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

func banner() {
	fmt.Println("")
	fmt.Println("[Author] marco.milon@gmail.com")
	fmt.Println(" ______     _____  _    _ _____  ")
	fmt.Println("|  ____|   |  __ \\| |  | |  __ \\ ")
	fmt.Println("| |__   ___| |__) | |__| | |__) |")
	fmt.Println("|  __| |_  /  ___/|  __  |  ___/ ")
	fmt.Println("| |____ / /| |    | |  | | |     ")
	fmt.Println("|______/___|_|    |_|  |_|_|     ")
	fmt.Println("")
}
