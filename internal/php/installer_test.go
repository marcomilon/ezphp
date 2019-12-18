package php_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/marcomilon/ezphp/internal/php"
)

var (
	temp        = os.TempDir() + "ezphptest"
	zipfile     = "out.zip"
	destination = temp + string(os.PathSeparator) + "out"
	files       = []struct {
		name, body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
)

func TestInstall(t *testing.T) {

	ioCom := setupClient()
	ts := setupInstall(t)
	defer ts.Close()

	installer := php.Installer{
		Source:      ts.URL + "/" + zipfile,
		Destination: temp + string(os.PathSeparator) + "out",
	}

	err := installer.Install(ioCom)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	for _, file := range files {
		path := destination + string(os.PathSeparator) + file.name
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected %v; got %v", nil, err)
		}
	}

	teardown(t)
}

func setupClient() php.IOCom {
	ioCom := php.IOCom{
		make(chan string),
		make(chan string),
		make(chan string),
		make(chan bool),
	}

	go func(ioCom php.IOCom) {
	Terminal:
		for {
			select {
			case outmsg := <-ioCom.Stdout:
				fmt.Fprintf(ioutil.Discard, outmsg)
			case <-ioCom.Done:
				break Terminal
			}
		}
	}(ioCom)

	return ioCom
}

func setupInstall(t *testing.T) *httptest.Server {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)

		// Create a new zip archive.
		zipWriter := zip.NewWriter(buf)

		// Add some files to the archive.
		for _, file := range files {
			f, err := zipWriter.Create(file.name)
			if err != nil {
				t.Fatal(err)
			}
			_, err = f.Write([]byte(file.body))
			if err != nil {
				t.Fatal(err)
			}
		}

		// Make sure to check the error on Close.
		err := zipWriter.Close()
		if err != nil {
			t.Fatal(err)
		}

		w.Header().Set("Content-type", "application/x-zip-compressed")
		w.Header().Set("Content-Disposition", "attachment; filename=out.zip")
		_, err = buf.WriteTo(w)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Fprintf(w, "%s", err)
	}))

	return ts
}

func teardown(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(temp, "*"))
	if err != nil {
		t.Fatal("Unable to destroy test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to destroy test")
		}
	}

	os.Remove(temp)
}
