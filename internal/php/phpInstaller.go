package php

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

type WriteCounter struct {
	total  int64
	length int64
	ioCom  IOCom
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.total += int64(n)

	percentage := float64(wc.total) / float64(wc.length) * float64(100)

	wc.ioCom.Stdout <- fmt.Sprintf("\rDownload in progress: %.2f%%", percentage)
	return n, nil
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

	err = i.download(ioCom)
	if err != nil {
		ioCom.Stderr <- "Download error: " + err.Error() + "\n"
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

func (i PhpInstaller) download(ioCom IOCom) error {

	if _, err := os.Stat(i.installDir); os.IsNotExist(err) {
		err = os.MkdirAll(i.installDir, 0755)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(i.installDir + string(os.PathSeparator) + i.filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(i.downloadUrl + "/" + i.filename)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{
		0,
		resp.ContentLength,
		ioCom,
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	ioCom.Stdout <- "\n"

	return nil
}

func (i PhpInstaller) unzip() error {

	src := i.installDir + string(os.PathSeparator) + i.filename
	dest := i.installDir

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

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
