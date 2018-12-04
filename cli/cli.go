package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/ezargs"
	"github.com/marcomilon/ezphp/engine/ezio"
	"github.com/marcomilon/ezphp/engine/fs"
	"github.com/marcomilon/ezphp/engine/php"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives"
	fileName     = "php-7.0.0-Win32-VC14-x64.zip"
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

type serveArguments struct {
	address      string
	documentRoot string
}

func Start(args ezargs.Arguments) {

	var ezIO ezio.EzIO = WhiteIO{}
	var phpPath string
	var err error

	if args.About {
		about()
		os.Exit(0)
	}

	var installer php.EzInstaller = php.Installer{
		downloadUrl,
		fileName,
		args.InstallDir,
	}

	ezIO.Info("EzPHP v" + php.EzPHPVersion + "\n")
	ezIO.Info("Website: " + php.EzPHPWebsite + "\n")
	ezIO.Info("\n")

	phpPath, err = fs.WhereIsPHP(args.InstallDir)
	if err != nil {
		localPHP, _ := filepath.Abs(args.InstallDir)
		ezIO.Info("Installing PHP v7.0.0 in local directory: " + localPHP + "\n")
		ezIO.Info("Please wait ... ")
		err = installer.Install(ezIO)
		phpPath = localPHP + string(os.PathSeparator) + php.PHP_EXECUTABLE
		ezIO.Info("\nPHP Installed succefully\n")
	}

	fs.CreateDirIfNotExist(args.DocRoot)

	localDocRoot, _ := filepath.Abs(args.DocRoot)

	ezIO.Info("Copy for PHP file in the directory: " + localDocRoot + "\n")
	ezIO.Info("Open your browser to http://" + args.Host + "\n")
	ezIO.Info("Web server is running...\n")
	ezIO.Info("Press CTRL+C to exit\n")

	phpServer := php.Server{
		phpPath,
		args.Host,
		args.DocRoot,
	}

	phpServer.Serve(ezIO, ezIO)

	byebye(ezIO)
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
	fmt.Printf("website: %s\n", php.EzPHPWebsite)
	fmt.Printf("version: %s\n", php.EzPHPVersion)
}
