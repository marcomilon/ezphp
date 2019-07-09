package app

import (
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/fs"
	"github.com/marcomilon/ezphp/engine/php"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives"
	fileName     = "php-7.0.0-Win32-VC14-x64.zip"
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
	installDir   = "php/7.0.0"
)

func Start(phpInstaller php.Installer, phpServer php.Server, ioCom php.IOCom, arguments php.Arguments) {

	var err error
	var phpExe string

	ioCom.Stdout <- "EzPHP v" + ezPHPVersion + "\n"
	ioCom.Stdout <- ezPHPWebsite + "\n"
	ioCom.Stdout <- "\n"

	phpExe, err = fs.WhereIsPHP(installDir)
	if err != nil {

		ioCom.Stdin <- "Would you like to install PHP?"
		result := <-ioCom.Stdin

		if result == "No" {
			ioCom.Done <- true
		}

		phpInstaller.Install(ioCom)

		ioCom.Stdout <- "\nPHP Installed succefully\n\n"

	}

	fs.CreateDirIfNotExist(arguments.DocRoot)

	localDocRoot, _ := filepath.Abs(arguments.DocRoot)

	ioCom.Stdout <- "Information\n"
	ioCom.Stdout <- "-----------\n"
	ioCom.Stdout <- "Copy your files to: " + localDocRoot + "\n"
	ioCom.Stdout <- "Web Server is available at http://" + arguments.Host + "/\n"
	ioCom.Stdout <- "Web server is running ...\n"

	phpServer.Serve(phpExe, ioCom)

	ioCom.Done <- true
}
