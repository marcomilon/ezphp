package php

import (
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
	resp, err := grab.Get(i.InstallDir+string(os.PathSeparator)+i.Filename, i.DownloadUrl+"/"+i.Filename)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i Installer) unzip() error {
	archiver.Unarchive(i.InstallDir+string(os.PathSeparator)+i.Filename, i.InstallDir)

	return nil
}
