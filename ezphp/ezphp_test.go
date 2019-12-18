package ezphp_test

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

	"github.com/marcomilon/ezphp/ezphp"
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

func TestPHPNotInPath(t *testing.T) {

	ioCom := php.IOCom{
		Stdout: make(chan string),
		Stderr: make(chan string),
		Stdin:  make(chan string),
		Done:   make(chan bool),
	}

	instl := php.Installer{
		Source:      "",
		Destination: "",
	}

	srv := php.Server{
		Exec:    "php7/notfound.exe",
		Host:    "",
		DocRoot: "",
	}

	go ezphp.Start(srv, instl, ioCom)

	in := <-ioCom.Stdin
	expected := "Would you like to install PHP?"

	if in != expected {
		t.Errorf("expected %v; got %v", expected, in)
	}

	ioCom.Stdin <- "No"
	done := <-ioCom.Done
	if !done {
		t.Errorf("expected %v; got %v", true, done)
	}

}

func TestInstallPHP(t *testing.T) {

	ts := setupInstall(t)
	defer ts.Close()

	ioCom := php.IOCom{
		Stdout: make(chan string),
		Stderr: make(chan string),
		Stdin:  make(chan string),
		Done:   make(chan bool),
	}

	instl := php.Installer{
		Source:      ts.URL + "/" + zipfile,
		Destination: temp + string(os.PathSeparator) + "out",
	}

	srv := php.Server{
		Exec:    "php7/notfound.exe",
		Host:    "",
		DocRoot: "",
	}

	go ezphp.Start(srv, instl, ioCom)

	in := <-ioCom.Stdin
	expected := "Would you like to install PHP?"

	if in != expected {
		t.Errorf("expected %v; got %v", expected, in)
	}

	ioCom.Stdin <- "Yes"

Installer:
	for {
		outmsg := <-ioCom.Stdout
		if outmsg == "PHP Installed succefully" {
			break Installer
		}
		fmt.Fprintf(ioutil.Discard, outmsg)
	}

	teardown(t)

}

func TestCreateDocumentRoot(t *testing.T) {

	ts := setupInstall(t)
	defer ts.Close()

	ioCom := php.IOCom{
		Stdout: make(chan string),
		Stderr: make(chan string),
		Stdin:  make(chan string),
		Done:   make(chan bool),
	}

	instl := php.Installer{
		Source:      ts.URL + "/" + zipfile,
		Destination: temp + string(os.PathSeparator) + "out",
	}

	docroot := temp + string(os.PathSeparator) + "docroot"

	srv := php.Server{
		Exec:    "php7/notfound.exe",
		Host:    "",
		DocRoot: docroot,
	}

	go ezphp.Start(srv, instl, ioCom)

	in := <-ioCom.Stdin
	expected := "Would you like to install PHP?"

	if in != expected {
		t.Errorf("expected %v; got %v", expected, in)
	}

	ioCom.Stdin <- "Yes"

Installer:
	for {
		outmsg := <-ioCom.Stdout
		if outmsg == "PHP Installed succefully" {
			break Installer
		}
		fmt.Fprintf(ioutil.Discard, outmsg)
	}

	if _, err := os.Stat(docroot); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

	teardown(t)

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
