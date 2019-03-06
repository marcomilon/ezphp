package php

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
	"github.com/sirupsen/logrus"
)

type Installer struct {
	DownloadUrl string
	Filename    string
	InstallDir  string
}

func (i *Installer) InstallPHP(ioCom IOCom) {

	var err error

	_, err = i.download(ioCom)
	if err != nil {
		logrus.Error("Error downloading file " + err.Error())
		ioCom.Errmsg <- err.Error()
		return
	}

	err = i.unzip()
	if err != nil {
		logrus.Error("Error unzipping file " + err.Error())
		ioCom.Errmsg <- err.Error()
		return
	}

}

func (i Installer) download(ioCom IOCom) (*grab.Response, error) {
	logrus.Info("Downloading PHP from " + i.DownloadUrl + "/" + i.Filename)
	client := grab.NewClient()
	req, _ := grab.NewRequest(i.InstallDir+string(os.PathSeparator)+i.Filename, i.DownloadUrl+"/"+i.Filename)
	resp := client.Do(req)
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			ioCom.Outmsg <- fmt.Sprintf("\rDownload in progress: %.2f%%", 100*resp.Progress())

		case <-resp.Done:
			break Loop
		}

	}

	ioCom.Outmsg <- fmt.Sprint("\rDownload in progress: 100%  ")

	if err := resp.Err(); err != nil {
		return nil, resp.Err()
	}

	return resp, nil
}

func (i Installer) unzip() error {
	logrus.Info("Unziping local PHP installation: " + i.InstallDir + string(os.PathSeparator) + i.Filename)
	err := archiver.Unarchive(i.InstallDir+string(os.PathSeparator)+i.Filename, i.InstallDir)
	if err != nil {
		return err
	}

	return nil
}
