package cli

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

func Start() {

	var (
		defaultExecPath string
		err             error
		pathToPHP       string
		ezOut           io.Writer
		ezErr           io.Writer
	)

	phpExec := flag.String("php", php.PHP_EXECUTABLE, "Path to php executable")
	serve := flag.String("S", "localhost:8080", "Listening address: <addr>:<port> ")
	documentRoot := flag.String("t", ".", "Document root: default value current directory")

	flag.Parse()

	defaultExecPath, err = fs.WhereIsGlobalPHP(*phpExec)
	if err != nil {

		ezio.Error(fmt.Sprintf("%s\n", err.Error()))

		defaultExecPath, err = fs.WhereIsLocalPHP(*phpExec, target)

		if err != nil {
			ezio.Error(fmt.Sprintf("%s\n", err.Error()))
			if runtime.GOOS == "windows" {
				if ezio.Confirm("Would you like to install PHP locally") {

					absTargetPath, _ := filepath.Abs(filepath.Dir(target))
					ezio.Info(fmt.Sprintf("Downloading PHP from %s\n", downloadUrl))
					ezio.Info(fmt.Sprintf("Installing PHP in %s\n", absTargetPath))

					pathToPHP, err = php.DownloadAndInstallPHP(downloadUrl, version, target)

					if err != nil {
						ezio.Error("Something went wrong\n")
						ezio.Error(fmt.Sprintf("%s\n", err.Error()))
						bybye()
					}

					defaultExecPath = pathToPHP + php.PHP_EXECUTABLE
				} else {
					bybye()
				}
			} else {
				ezio.Info(fmt.Sprintf("%s: %s\n", "Installer not available in your Operation System", runtime.GOOS))
				ezio.Info("Please install PHP using your favorite package manager\n")
				bybye()
			}
		} else {
			ezio.Info(fmt.Sprintf("Local installation of PHP founded in: %s\n", defaultExecPath))
		}

	}

	if len(*serve) > 0 {
		absDocumentRootPath, _ := filepath.Abs(filepath.Dir(*documentRoot))
		ezio.Info("EzPHP\n")
		ezio.Info(fmt.Sprintf("website: %s\n", ezPHPWebsite))
		ezio.Info(fmt.Sprintf("Running PHP from: %s\n", defaultExecPath))
		ezio.Info("Server is ready\n")
		ezio.Info(fmt.Sprintf("Document root is: %s\n", absDocumentRootPath))
		ezio.Info(fmt.Sprintf("Open your web browser to: http://%s\n", *serve))

		ezOut = ezio.EzOut{Prompt: " Serve"}
		ezOut = ezOut
		ezErr = ezOut

	} else {
		ezOut = os.Stdout
		ezErr = os.Stderr
	}
	
	err = php.Run(defaultExecPath, os.Args, ezOut, ezErr)
	if err != nil {
		ezio.Error("Something went wrong\n")
		ezio.Error(fmt.Sprintf("%s\n", err.Error()))
		bybye()
	}

}

func bybye() {
	ezio.Info("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
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