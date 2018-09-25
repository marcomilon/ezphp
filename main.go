package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/marcomilon/ezphp/internals/helpers/ezio"
	"github.com/marcomilon/ezphp/internals/helpers/fs"
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

	var (
		versionFlag     bool
		defaultExecPath string
		err             error
		pathToPHP       string
	)

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
			ezio.Info(php.PHP_EXECUTABLE + " not found\n")
			if runtime.GOOS == "windows" {
				if ezio.Confirm("Would you like to install PHP locally") {
					pathToPHP, err = php.DownloadAndInstallPHP(downloadUrl, version, target)
					defaultExecPath = pathToPHP + php.PHP_EXECUTABLE
				} else {
					bybye()
					return
				}
			} else {
				ezio.Info("Auto installer not available in your Operation System\n")
				ezio.Info("Please install PHP using your favorite package manager\n")
				bybye()
				return
			}
		}

	}

	pathToDocRoot, _ := filepath.Abs(filepath.Dir(*public))
	ezio.Info(fmt.Sprintf("Running PHP from: %s\n", defaultExecPath))
	ezio.Info(fmt.Sprintf("Document root is: %s\n", pathToDocRoot))
	ezio.Info(fmt.Sprintf("Open your web browser to: http://%s\n", *host))
	err = php.Serve(defaultExecPath, *host, *public)
	if err != nil {
		ezio.Info(err.Error())
		bybye()
	}

}

func bybye() {
	ezio.Info("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
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
