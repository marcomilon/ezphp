package php

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
	"github.com/sirupsen/logrus"
)

const (
	downloadUrl = "https://windows.php.net/downloads/releases/archives"
	fileName    = "php-7.0.0-Win32-VC14-x64.zip"
	installDir  = "php/7.0.0"
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

func (i PhpInstaller) Install(ioCom IOCom) {

	var err error

	absPath, _ := filepath.Abs(filepath.Dir(i.installDir))
	localinstallDir := absPath + string(os.PathSeparator) + i.installDir

	ioCom.Stdout <- "\nInstalling PHP v7.0.0 in your local directory: " + localinstallDir + "\n"
	ioCom.Stdout <- "Downloading PHP from: " + i.downloadUrl + "/" + i.filename + "\n"

	_, err = i.download(ioCom)
	if err != nil {
		logrus.Error("Error downloading file " + err.Error())
		ioCom.Stderr <- err.Error()
		return
	}

	err = i.unzip()
	if err != nil {
		logrus.Error("Error unzipping file " + err.Error())
		ioCom.Stderr <- err.Error()
		return
	}

}

func (i PhpInstaller) download(ioCom IOCom) (*grab.Response, error) {
	logrus.Info("Downloading PHP from " + i.downloadUrl + "/" + i.filename)
	client := grab.NewClient()
	req, _ := grab.NewRequest(i.installDir+string(os.PathSeparator)+i.filename, i.downloadUrl+"/"+i.filename)
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

	ioCom.Stdout <- fmt.Sprint("\rDownload in progress: 100%  ")

	if err := resp.Err(); err != nil {
		return nil, resp.Err()
	}

	return resp, nil
}

func (i PhpInstaller) unzip() error {
	logrus.Info("Unziping local PHP installation: " + i.installDir + string(os.PathSeparator) + i.filename)
	err := archiver.Unarchive(i.installDir+string(os.PathSeparator)+i.filename, i.installDir)
	if err != nil {
		return err
	}

	return nil
}
