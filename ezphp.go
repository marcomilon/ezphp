package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/marcomilon/ezphp/internals/install"
	"github.com/marcomilon/ezphp/internals/output"
	"github.com/marcomilon/ezphp/internals/serve"
)

const (
	phpExe = "php.exe"
)

func main() {

	var defaultExecPath string
	var err error

	banner()

	php := flag.String("php", "", "Path to php executable")
	host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
	public := flag.String("public", "public", "Path to public directory")

	flag.Parse()

	if *php == "" {

		defaultExecPath, err = searchForPhp(phpExe)
		if err != nil {
			output.Error(err.Error() + "\n")
			output.Info("Press Enter to exit... ")
			fmt.Scanln()
			fmt.Scanln()
			return
		}

		php = &defaultExecPath
	}

	output.Info("Your document root directory is: " + *public + "\n")
	install.CreateDirIfNotExist(*public)
	
	err = serve.Start(*php, *host, *public)
	if err != nil {
		output.Error("Unable to execute PHP: " + err.Error() + "\n")
		output.Error("Press Enter to continue... ")
		fmt.Scanln()
	}

	return
}

func searchForPhp(phpExe string) (string, error) {

	var defaultExecPath string
	var path string
	var err error
	var absPath string

	output.Info("Looking for php in default directory: " + install.PhpDir + "\n")
	if _, err = os.Stat(install.PhpDir); err == nil {
		output.Info("Local php installation founded\n")
		absPath, _ = filepath.Abs(filepath.Dir(install.PhpDir))
		defaultExecPath = absPath + string(os.PathSeparator) + install.PhpDir + string(os.PathSeparator) + phpExe
		return defaultExecPath, nil
	}

	defaultExecPath, err = exec.LookPath(phpExe)
	if err != nil {
		output.Error("php executable not found in path\n")

		if !askToInstallPhp() {
			return "", errors.New("php won't be installed. bye bye.")
		}

		path, err = install.Installer(install.Version, install.PhpDir)
		if err != nil {
			return "", errors.New(err.Error())
		}

		defaultExecPath = path + string(os.PathSeparator) + phpExe
	}

	return defaultExecPath, nil
}

func askToInstallPhp() bool {
	var confirmation string

	output.Installer("Would you like to install php locally [y/N]? ")
	fmt.Scan(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" || confirmation == "yes" {
		return true
	}

	return false
}

func banner() {
	fmt.Println(" ______     _____  _    _ _____  ")
	fmt.Println("|  ____|   |  __ \\| |  | |  __ \\ ")
	fmt.Println("| |__   ___| |__) | |__| | |__) |")
	fmt.Println("|  __| |_  /  ___/|  __  |  ___/ ")
	fmt.Println("| |____ / /| |    | |  | | |     ")
	fmt.Println("|______/___|_|    |_|  |_|_|     ")
	fmt.Println("Author", "marco.milon@gmail.com")
	fmt.Println("")
}
