package php

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
)

const (
	downloadUrl = "https://windows.php.net/downloads/releases/archives"
	fileName    = "php-7.0.0-Win32-VC14-x64.zip"
	installDir  = "php/7.0.0"
	version     = "7.0.0"
)

type PhpInstaller struct {
	downloadUrl string
	filename    string
	installDir  string
}

func NewPhpInstaller() PhpInstaller {
	return PhpInstaller{
		downloadUrl,
		fileName,
		installDir,
	}
}

func (i PhpInstaller) Install(ioCom IOCom) (string, error) {

	var err error

	localinstallDir, _ := filepath.Abs(filepath.Dir(i.installDir))

	ioCom.Stdout <- "Downloading PHP from: " + i.downloadUrl + "/" + i.filename + "\n"
	ioCom.Stdout <- "Please wait...\n"

	_, err = i.download(ioCom)
	if err != nil {
		ioCom.Stderr <- "Error: " + err.Error() + "\n"
		return "", err
	}

	ioCom.Stdout <- "Installing PHP v" + version + " in your local directory: " + localinstallDir + "\n"

	err = i.unzip()
	if err != nil {
		ioCom.Stderr <- "Error: " + err.Error() + "\n"
		return "", err
	}

	return localinstallDir + string(os.PathSeparator) + version + string(os.PathSeparator) + PHP_EXECUTABLE, nil

}

func (i PhpInstaller) download(ioCom IOCom) (*grab.Response, error) {
	client := grab.NewClient()
	req, err := grab.NewRequest(i.installDir+string(os.PathSeparator)+i.filename, i.downloadUrl+"/"+i.filename)
	if err != nil {
		return nil, err
	}

	resp := client.Do(req)
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			ioCom.Stdout <- fmt.Sprintf("\rDownload in progress: %.2f%%", 100*resp.Progress())

		case <-resp.Done:
			break Loop
		}

	}

	if err := resp.Err(); err != nil {
		return nil, resp.Err()
	}

	ioCom.Stdout <- fmt.Sprint("\rDownload in progress: 100%  \n")

	return resp, nil
}

func (i PhpInstaller) unzip() error {
	return archiver.Unarchive(i.installDir+string(os.PathSeparator)+i.filename, i.installDir)
}
