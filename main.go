package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/marcomilon/ezphp/internals/helpers/fs"
	"github.com/marcomilon/ezphp/internals/helpers/output"
	"github.com/marcomilon/ezphp/internals/helpers/prompt"
	"github.com/marcomilon/ezphp/internals/php"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives/"
	version      = "php-7.0.0-Win32-VC14-x64.zip"
	target       = "php-7.0.0"
	ezPHPVersion = "0.0.1"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

func main() {

	var versionFlag bool
	var defaultExecPath string
	var err error

	phpExec := flag.String("php", php.PHP_EXECUTABLE, "Path to php executable")
	host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
	public := flag.String("public", ".", "Path to public directory")
	flag.BoolVar(&versionFlag, "v", false, "Prints about message")

	flag.Parse()

	if versionFlag {
		about()
		return
	}

	defaultExecPath, err = fs.WhereIsGlobalPHP(*phpExec)
	if err != nil {

		defaultExecPath, err = fs.WhereIsLocalPHP(*phpExec)

		if err != nil {
			output.Info("PHP not installed\n")
			if runtime.GOOS == "linux" {
				if prompt.Confirm("Would you like to install PHP locally") {
					defaultExecPath, err = php.DownloadAndInstallPHP(downloadUrl, version, target)
				}
			} else {
				output.Info("Auto installer not available in your Operation System\n")
				output.Info("Please install PHP using your favorite package manager\n")
			}
		}

	}

	php.Serve(defaultExecPath, *host, *public)

}

func about() {
	fmt.Println(" ______     _____  _    _ _____  ")
	fmt.Println("|  ____|   |  __ \\| |  | |  __ \\ ")
	fmt.Println("| |__   ___| |__) | |__| | |__) |")
	fmt.Println("|  __| |_  /  ___/|  __  |  ___/ ")
	fmt.Println("| |____ / /| |    | |  | | |     ")
	fmt.Println("|______/___|_|    |_|  |_|_|     ")
	fmt.Println("")
	fmt.Printf("website: %s\n", ezPHPWebsite)
	fmt.Printf("version: %s\n", ezPHPVersion)
}
