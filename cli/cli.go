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
	"github.com/sirupsen/logrus"
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
	ezIO.Info("PHP version: " + php.PHPVersion + "\n")
	absDoctRoot, _ := filepath.Abs(args.DocRoot)
	ezIO.Info("Your web server document root is: " + absDoctRoot + "\n")
	ezIO.Info("\n")

	phpPath, err = fs.WhereIsPHP(args.InstallDir)
	if err != nil {
		logrus.Info("PHP not founded")

		localPHP, _ := filepath.Abs(args.InstallDir)

		if ezIO.Confirm("Would you like to install PHP locally?") {

			ezIO.Info("Installing PHP v7.0.0 in your local directory: " + localPHP + "\n")
			ezIO.Info("Downloading PHP from: " + downloadUrl + "/" + fileName + "\n")
			ezIO.Info("File size is ~24MB\n")
			ezIO.Info("This may take a while\n")
			err = installer.Install(ezIO)

			if err != nil {
				ezIO.Error("\nFailed to install PHP: " + err.Error() + "\n")
				byebye(ezIO)
			}

			phpPath = localPHP + string(os.PathSeparator) + php.PHP_EXECUTABLE
			ezIO.Info("\nPHP Installed succefully\n\n")

		} else {
			byebye(ezIO)
		}
	}

	fs.CreateDirIfNotExist(args.DocRoot)

	localDocRoot, _ := filepath.Abs(args.DocRoot)

	ezIO.Info("Information\n")
	ezIO.Info("-----------\n")
	ezIO.Info("Copy for PHP file in the directory: " + localDocRoot + "\n")
	ezIO.Info("Open your browser to http://" + args.Host + "\n")
	ezIO.Info("Web server is running ...\n")
	ezIO.Info("Press CTRL+C to exit\n")

	phpServer := php.Server{
		phpPath,
		args.Host,
		args.DocRoot,
	}

	err = phpServer.Serve(ezIO, ezIO)
	if err != nil {
		ezIO.Error(err.Error() + "\n")
	}

	ezIO.Info("Failed to start web server. See error above ^\n")
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
