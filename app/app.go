package app

import (
	"os"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/ezargs"
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

type buffer struct {
	php.IOCom
}

func Start(args ezargs.Arguments, ioChannels php.IOCom) {

	var phpPath string
	var err error

	installer := php.Installer{
		downloadUrl,
		fileName,
		installDir,
	}

	buffer := buffer{
		ioChannels,
	}

	buffer.out("EzPHP v" + ezPHPVersion + "\n")
	buffer.out(ezPHPWebsite + "\n")
	buffer.out("\n")

	phpPath, err = fs.WhereIsPHP(installDir)
	if err != nil {

		ioChannels.Confirm <- "Would you like to install PHP?"
		result := <-ioChannels.Confirm

		if result == "No" {
			ioChannels.Done <- true
		}

		logrus.Info("PHP not founded")

		localPHP, _ := filepath.Abs(installDir)

		installer.InstallPHP(ioChannels)

		phpPath = localPHP + string(os.PathSeparator) + php.PHP_EXECUTABLE
		buffer.out("\nPHP Installed succefully\n\n")

	}

	fs.CreateDirIfNotExist(baseDir + string(os.PathSeparator) + args.DocRoot)

	localDocRoot, _ := filepath.Abs(baseDir + string(os.PathSeparator) + args.DocRoot)

	buffer.out("Information\n")
	buffer.out("-----------\n")
	buffer.out("Copy your files to: " + localDocRoot + "\n")
	buffer.out("Web Server is available at http://" + args.Host + "/\n")
	buffer.out("Web server is running ...\n")

	phpServer := php.Server{
		phpPath,
		args.Host,
		baseDir + string(os.PathSeparator) + args.DocRoot,
	}

	phpServer.StartServer(ioChannels)

	ioChannels.Done <- true
}

func (b buffer) out(msg string) {
	outmsg := php.NewStdout(msg)
	b.Outmsg <- outmsg
}
