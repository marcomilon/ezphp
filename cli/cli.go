package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/marcomilon/ezphp/engine/ezio"
	"github.com/marcomilon/ezphp/engine/php"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives/"
	version      = "php-7.0.0-Win32-VC14-x64.zip"
	destination  = "php-7.0.0/"
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

type serveArguments struct {
	address      string
	documentRoot string
}

func Clean() {

	var installer php.EzInstaller = php.Installer{
		downloadUrl,
		destination,
		version,
	}

	var ezIO ezio.EzIO = InOut{}

	phpPath, err := installer.WhereIs()
	if err != nil {
		ezIO.Info("Installing...\n")
		installer.Install(ezIO)
	}

	ezIO.Info(fmt.Sprintf("php found ... %s\n", phpPath))

	phpServer := php.Server{
		phpPath,
		".",
		"localhost",
		"8080",
	}

	phpServer.Serve(ezIO, ezIO)

	byebye(ezIO)
}

func Start() {

	// var (
	// 	defaultExecPath string
	// 	err             error
	// 	pathToPHP       string
	// )
	//
	// phpExec := php.PHP_EXECUTABLE
	// ezIO := InOut{}
	//
	// defaultExecPath, err = fs.WhereIsGlobalPHP(phpExec)
	//
	// if err != nil {
	//
	// 	ezIO.Error(fmt.Sprintf("%s\n", err.Error()))
	//
	// 	defaultExecPath, err = fs.WhereIsLocalPHP(phpExec, target)
	//
	// 	if err != nil {
	//
	// 		ezIO.Error(fmt.Sprintf("%s\n", err.Error()))
	//
	// 		if runtime.GOOS == "windows" {
	//
	// 			if ezIO.Confirm("Would you like to install PHP locally") {
	//
	// 				pathToPHP, err = php.DownloadAndInstallPHP(downloadUrl, version, target, ezIO)
	//
	// 				if err != nil {
	// 					ezIO.Error("Something went wrong\n")
	// 					ezIO.Error(fmt.Sprintf("%s\n", err.Error()))
	// 					byebye(ezIO)
	// 				}
	//
	// 				defaultExecPath = pathToPHP + php.PHP_EXECUTABLE
	// 			} else {
	// 				byebye(ezIO)
	// 			}
	//
	// 		} else {
	// 			ezIO.Info(fmt.Sprintf("%s: %s\n", "Installer not available in your Operation System", runtime.GOOS))
	// 			ezIO.Info("Please install PHP using your favorite package manager\n")
	// 			byebye(ezIO)
	// 		}
	// 	} else {
	// 		ezIO.Info(fmt.Sprintf("Local installation of PHP founded in: %s\n", defaultExecPath))
	// 	}
	// }
	//
	// serverArguments := parseServeFlag()
	// absDocumentRootPath, _ := filepath.Abs(filepath.Dir(serverArguments.documentRoot))
	// fs.CreateDirIfNotExist(fmt.Sprintf("%s/%s", absDocumentRootPath, serverArguments.documentRoot))
	//
	// ezIO.Info("EzPHP\n")
	// ezIO.Info(fmt.Sprintf("website: %s\n", ezPHPWebsite))
	// ezIO.Info(fmt.Sprintf("Running PHP from: %s\n", defaultExecPath))
	// ezIO.Info("Server is ready\n")
	// ezIO.Info(fmt.Sprintf("Document root is: %s/%s\n", absDocumentRootPath, serverArguments.documentRoot))
	// ezIO.Info(fmt.Sprintf("Open your web browser to: http://%s\n", serverArguments.address))
	//
	// err = php.Run(defaultExecPath, os.Args, ezIO, ezIO)
	//
	// if err != nil {
	// 	ezIO.Error("Something went wrong\n")
	// 	ezIO.Error(fmt.Sprintf("%s\n", err.Error()))
	// 	byebye(ezIO)
	// }

}

func byebye(ezIO ezio.EzIO) {
	ezIO.Info("Press 'Enter' to exit...")
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

func startWebServer() bool {
	return contains(os.Args, "-S")
}

func parseServeFlag() serveArguments {
	address := flag.String("S", "localhost:8080", "-S <addr>:<port>\tRun with built-in web server.")
	documentRoot := flag.String("t", "public_html", "-t <docroot>\t\tSpecify document root <docroot> for built-in web server.")

	flag.Parse()

	args := serveArguments{
		address:      *address,
		documentRoot: *documentRoot,
	}

	return args
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
