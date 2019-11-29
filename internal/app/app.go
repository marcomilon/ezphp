package app

import (
	"path/filepath"

	"github.com/marcomilon/ezphp/internal/fs"
	"github.com/marcomilon/ezphp/internal/php"
)

const (
	installDir = "php/7.0.0"
)

func StartServer(phpInstaller php.Installer, phpServer php.Server, ioCom php.IOCom, arguments php.Arguments) {

	var err error
	var phpExe string

	phpExe, err = fs.WhereIsPHP(installDir)

	if err != nil {

		ioCom.Stdin <- "Would you like to install PHP?"
		result := <-ioCom.Stdin

		if result == "No" {
			ioCom.Done <- true
		}

		phpExe, err = phpInstaller.Install(ioCom)
		if err != nil {
			ioCom.Stdout <- "Unable to install PHP\n"
			ioCom.Done <- true
		}

		ioCom.Stdout <- "PHP Installed succefully\n\n"

	}

	fs.CreateDirIfNotExist(arguments.DocRoot)

	localDocRoot, _ := filepath.Abs(arguments.DocRoot)

	ioCom.Stdout <- "Information\n"
	ioCom.Stdout <- "-----------\n"
	ioCom.Stdout <- "PHP found in " + phpExe + "\n"
	ioCom.Stdout <- "Copy your .php files to: " + localDocRoot + "\n"
	ioCom.Stdout <- "Web Server is available at http://" + arguments.Host + "/\n"
	ioCom.Stdout <- "Web server is running ...\n"

	phpServer.Serve(phpExe, ioCom)

	ioCom.Done <- true
}
