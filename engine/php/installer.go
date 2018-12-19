package php

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
	"github.com/sirupsen/logrus"
)

func (i *Installer) Install() {

	var err error

	_, err = i.download()
	if err != nil {
		i.Errmsg <- err.Error()
		return
	}

	err = i.unzip()
	if err != nil {

	}

	i.Done <- true

}

func (i Installer) download() (*grab.Response, error) {
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
			// fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
			// 	resp.BytesComplete(),
			// 	resp.Size,
			// 	100*resp.Progress())
			i.Outmsg <- fmt.Sprintf("%.2f%%", 100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}

	}

	if err := resp.Err(); err != nil {
		logrus.Info("Download error")
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
