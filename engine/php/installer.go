package php

import (
	"os"

	"github.com/cavaliercoder/grab"
	"github.com/marcomilon/ezphp/engine/ezio"
)

func (i Installer) Install(w ezio.EzIO) error {
	i.download()
	return nil
}

func (i Installer) download() (*grab.Response, error) {
	resp, err := grab.Get(i.InstallDir+string(os.PathSeparator)+i.Filename, i.DownloadUrl+"/"+i.Filename)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
