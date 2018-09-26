package php

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"github.com/marcomilon/ezphp/internals/helpers/ezio"
	"github.com/marcomilon/ezphp/internals/helpers/fs"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", "")

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	ezio.Custom(fmt.Sprintf("Please wait... %s complete", humanize.Bytes(wc.Total)))
}

func DownloadAndInstallPHP(downloadUrl string, version string, destination string) (string, error) {

	var (
		absPath string
		err     error
	)

	err = fs.CreateDirIfNotExist(destination)
	if err != nil {
		return "", err
	}

	err = download(downloadUrl+version, destination+string(os.PathSeparator)+version)
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

func download(url string, dest string) error {

	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Unable to download File: " + url)
	}

	// Create the file
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Print("\n")

	return nil
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
