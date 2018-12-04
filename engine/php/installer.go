package php

import (
	"log"
	"os"

	"github.com/cavaliercoder/grab"
	"github.com/marcomilon/ezphp/engine/ezio"
	"github.com/mholt/archiver"
)

func (i Installer) Install(w ezio.EzIO) error {
	i.download()
	i.unzip()
	return nil
}

func (i Installer) download() (*grab.Response, error) {
	log.Println("Downloading PHP from " + i.DownloadUrl + "/" + i.Filename)
	resp, err := grab.Get(i.InstallDir+string(os.PathSeparator)+i.Filename, i.DownloadUrl+"/"+i.Filename)
	if err != nil {
		log.Println("Failed to download PHP: " + err.Error())
		return nil, err
	}

	return resp, nil
}

func (i Installer) unzip() error {
	log.Println("Unziping local PHP installation: " + i.InstallDir + string(os.PathSeparator) + i.Filename)
	archiver.Unarchive(i.InstallDir+string(os.PathSeparator)+i.Filename, i.InstallDir)

	return nil
}
