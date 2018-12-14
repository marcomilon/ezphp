package php

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/marcomilon/ezphp/engine/ezio"
	"github.com/mholt/archiver"
	"github.com/sirupsen/logrus"
)

func (i Installer) Install(w ezio.EzIO) error {

	var err error

	_, err = i.download()
	if err != nil {
		return err
	}

	err = i.unzip()
	if err != nil {
		return err
	}

	return nil
}

func (i Installer) download() (*grab.Response, error) {
	logrus.Info("Downloading PHP from " + i.DownloadUrl + "/" + i.Filename)
	resp, err := grab.Get(i.InstallDir+string(os.PathSeparator)+i.Filename, i.DownloadUrl+"/"+i.Filename)
	if err != nil {
		logrus.Error("Failed to download PHP: " + err.Error())
		return nil, err
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

func spinner(delay time.Duration) {
	StopSpinner := false
	for !StopSpinner {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
