package app

import (
	"os"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/fs"
	"github.com/marcomilon/ezphp/engine/php"
	"github.com/sirupsen/logrus"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives"
	fileName     = "php-7.0.0-Win32-VC14-x64.zip"
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
	baseDir      = "ezphp"
	installDir   = "ezphp/php/7.0.0"
)

func Start(phpInstaller php.Installer, phpServer php.Server, ioCom php.IOCom) {

	var err error

	ioCom.Stdout <- "EzPHP v" + ezPHPVersion + "\n"
	ioCom.Stdout <- ezPHPWebsite + "\n"
	ioCom.Stdout <- "\n"

	_, err = fs.WhereIsPHP(installDir)
	if err != nil {

		ioCom.Stdin <- "Would you like to install PHP?"
		result := <-ioCom.Stdin

		if result == "No" {
			ioCom.Done <- true
		}

		logrus.Info("PHP not founded")

		phpInstaller.Install(ioCom)

		ioCom.Stdout <- "\nPHP Installed succefully\n\n"

	}

	fs.CreateDirIfNotExist(baseDir + string(os.PathSeparator) + phpServer.GetDocRoot())

	localDocRoot, _ := filepath.Abs(baseDir + string(os.PathSeparator) + phpServer.GetDocRoot())

	ioCom.Stdout <- "Information\n"
	ioCom.Stdout <- "-----------\n"
	ioCom.Stdout <- "Copy your files to: " + localDocRoot + "\n"
	ioCom.Stdout <- "Web Server is available at http://" + phpServer.GetHost() + "/\n"
	ioCom.Stdout <- "Web server is running ...\n"

	phpServer.Serve(ioCom)

	ioCom.Done <- true
}
