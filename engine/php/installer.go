package php

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/marcomilon/ezphp/engine/ezio"
	"github.com/marcomilon/ezphp/engine/fs"
)

func DownloadAndInstallPHP(downloadUrl string, version string, destination string, ezIO ezio.EzIO) (string, error) {

	var (
		absPath string
		err     error
	)

	err = fs.CreateDirIfNotExist(destination)
	if err != nil {
		return "", err
	}

	err = grabPHP(downloadUrl+version, destination+string(os.PathSeparator)+version, ezIO)
	if err != nil {
		return "", err
	}

	err = unzip(destination+string(os.PathSeparator)+version, destination)
	if err != nil {
		return "", err
	}

	absPath, err = filepath.Abs(filepath.Dir(destination))
	if err != nil {
		return "", err
	}

	path := absPath + string(os.PathSeparator) + destination + string(os.PathSeparator)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err != nil {
			return "", err
		}
	}

	return path, nil
}

func unzip(src string, dest string) error {

	r, err := zip.OpenReader(src)

	if err != nil {
		return err
	}

	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {

		rc, err := f.Open()

		if err != nil {
			return err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

			if err != nil {
				return err
			}

			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)

			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func grabPHP(url string, destination string, ezIO ezio.EzIO) error {
	client := grab.NewClient()
	req, _ := grab.NewRequest(destination, url)

	// start download
	ezIO.Info(fmt.Sprintf("Downloading %v...\n", req.URL()))
	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			ezIO.Custom("Please wait", fmt.Sprintf("Transferred %v / %v bytes (%.2f%%)\r",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress()))

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		ezIO.Error(fmt.Sprintf("Download failed: %v\n", err))
		return err
	}

	ezIO.Info(fmt.Sprintf("Download saved to ./%v \n", resp.Filename))
	return nil
}
