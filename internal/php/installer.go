package php

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type progressBar struct {
	total  int64
	length int64
}

func (pb *progressBar) Write(p []byte) (int, error) {
	n := len(p)
	pb.total += int64(n)

	percentage := float64(pb.total) / float64(pb.length) * float64(100)

	fmt.Printf("\rDownload in progress: %.2f%%", percentage)
	return n, nil
}

func FastInstall(source, installFolder string) (string, error) {

	var confirmation string

	fmt.Print("Would you like to install PHP version 7.4.14? [y/N] ")
	fmt.Scanln(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation != "y" {
		ExitEzPHP()
	}

	fmt.Printf("Downloading php from %v\n", source)
	fmt.Println("Please wait...")

	zipfile, err := download(source, installFolder)
	if err != nil {
		return "", err
	}

	fmt.Println("Installing PHP version 7.4.10")

	err = unzip(zipfile, installFolder)
	if err != nil {
		return "", err
	}

	abs, _ := filepath.Abs(installFolder)
	fmt.Printf("PHP was succefully installed in %v\n", abs)

	phpExe := fmt.Sprintf("%v/%v", installFolder, PHP_EXECUTABLE)

	return phpExe, nil
}

func download(source, installFolder string) (string, error) {

	filename := path.Base(source)

	resp, err := http.Get(source)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error %d\n", resp.StatusCode)
	}

	if _, err := os.Stat(installFolder); os.IsNotExist(err) {
		err = os.MkdirAll(installFolder, 0755)
		if err != nil {
			return "", err
		}
	}

	out, err := os.Create(installFolder + string(os.PathSeparator) + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	pb := &progressBar{
		0,
		resp.ContentLength,
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, pb))
	if err != nil {
		return "", err
	}

	fmt.Println("")

	zipfile := installFolder + string(os.PathSeparator) + filename

	return zipfile, nil
}

func unzip(source, installFolder string) error {

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

	os.MkdirAll(installFolder, 0755)

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

		path := filepath.Join(installFolder, f.Name)

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

func ExitEzPHP() {
	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}
