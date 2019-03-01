package app

import (
	"os"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine"
	"github.com/marcomilon/ezphp/engine/ezargs"
	"github.com/marcomilon/ezphp/engine/fs"
	"github.com/marcomilon/ezphp/engine/php"
	"github.com/marcomilon/ezphp/engine/php/install"
	"github.com/marcomilon/ezphp/engine/php/serve"
	"github.com/sirupsen/logrus"
)

const (
	downloadUrl  = "https://windows.php.net/downloads/releases/archives"
	fileName     = "php-7.0.0-Win32-VC14-x64.zip"
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

type buffer struct {
	engine.IOCom
}

func Start(args ezargs.Arguments, ioChannels engine.IOCom) {

	var phpPath string
	var err error

	installer := install.Installer{
		downloadUrl,
		fileName,
		args.InstallDir,
	}

	buffer := buffer{
		ioChannels,
	}

	buffer.out("EzPHP v" + ezPHPVersion + "\n")
	buffer.out(ezPHPWebsite + "\n")
	buffer.out("\n")

	phpPath, err = fs.WhereIsPHP(args.InstallDir)
	if err != nil {
		logrus.Info("PHP not founded")

		localPHP, _ := filepath.Abs(args.InstallDir)

		buffer.out("Installing PHP v7.0.0 in your local directory: " + localPHP + "\n")
		buffer.out("Downloading PHP from: " + downloadUrl + "/" + fileName + "\n")

		installer.InstallPHP(ioChannels)

		phpPath = localPHP + string(os.PathSeparator) + php.PHP_EXECUTABLE
		buffer.out("\nPHP Installed succefully\n\n")

	}

	fs.CreateDirIfNotExist(args.DocRoot)

	localDocRoot, _ := filepath.Abs(args.DocRoot)

	buffer.out("Information\n")
	buffer.out("-----------\n")
	buffer.out("Copy for PHP file in the directory: " + localDocRoot + "\n")
	buffer.out("Open your browser to http://" + args.Host + "\n")
	buffer.out("Web server is running ...\n")

	phpServer := serve.Server{
		phpPath,
		args.Host,
		args.DocRoot,
	}

	phpServer.StartServer(ioChannels)

	ioChannels.Done <- true
}

func (b buffer) out(msg string) {
	b.Outmsg <- msg
}
