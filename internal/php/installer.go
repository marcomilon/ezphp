package php

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Installer define the source url and the destionation folder
type Installer struct {
	Source      string
	Destination string
}

type progressBar struct {
	total  int64
	length int64
	ioCom  IOCom
}

func (pb *progressBar) Write(p []byte) (int, error) {
	n := len(p)
	pb.total += int64(n)

	percentage := float64(pb.total) / float64(pb.length) * float64(100)

	pb.ioCom.Stdout <- fmt.Sprintf("\rDownload in progress: %.2f%%", percentage)
	return n, nil
}

// Install tries to download php from source url and install it on the destination folder
func (i Installer) Install(ioCom IOCom) error {
	zipfile, err := download(i.Source, i.Destination, ioCom)
	if err != nil {
		return err
	}

	err = unzip(zipfile, i.Destination)
	if err != nil {
		return err
	}

	return nil
}

func unzip(source, destination string) error {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}

	defer func() error {
		if err := r.Close(); err != nil {
			return err
		}

		return nil
	}()

	os.MkdirAll(destination, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() error {
			if err := rc.Close(); err != nil {
				return err
			}

			return nil
		}()

		path := filepath.Join(destination, f.Name)

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

func download(source, destination string, ioCom IOCom) (string, error) {

	filename := path.Base(source)

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.MkdirAll(destination, 0755)
		if err != nil {
			return "", err
		}
	}

	resp, err := http.Get(source)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(destination + string(os.PathSeparator) + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	pb := &progressBar{
		0,
		resp.ContentLength,
		ioCom,
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, pb))
	if err != nil {
		return "", err
	}

	ioCom.Stdout <- "\n"

	zipfile := destination + string(os.PathSeparator) + filename

	return zipfile, nil
}
