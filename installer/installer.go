package installer

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	PhpExecutable  = "php.exe"
	PhpDir         = "php-7.0.0"
	phpDownloadUrl = "https://windows.php.net/downloads/releases/archives/php-7.0.0-Win32-VC14-x64.zip"
	phpZipFile     = "php-7.0.0-Win32-VC14-x64.zip"
)

var (
	absPath string
	path    string
	err     error
)

func Install() (string, error) {

	fmt.Println("[Installer] Installing PHP. Please wait...")

	err = CreateDirIfNotExist(PhpDir)
	if err != nil {
		return "", err
	}

	err = downloadFile(PhpDir + string(os.PathSeparator) + phpZipFile, phpDownloadUrl)
	if err != nil {
		return "", err
	}

	err = unzip(PhpDir + string(os.PathSeparator) + phpZipFile, PhpDir)
	if err != nil {
		return "", err
	}

	absPath, err = filepath.Abs(filepath.Dir(PhpDir))
	if err != nil {
		return "", err
	}

	path := absPath + string(os.PathSeparator) + PhpDir + string(os.PathSeparator) + PhpExecutable

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err != nil {
			return "", err
		}
	}

	return path, nil
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(filepath string, url string) error {

	if _, err := os.Stat(filepath); err == nil {
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
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func unzip(src, dest string) error {
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
