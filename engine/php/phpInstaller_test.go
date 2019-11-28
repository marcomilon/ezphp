package php

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var temp = os.TempDir() + "ezphptest"
var zipfile = "out.zip"
var output = temp + string(os.PathSeparator) + zipfile

var files = []struct {
	Name, Body string
}{
	{"readme.txt", "This archive contains some text files."},
	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	{"todo.txt", "Get animal handling licence.\nWrite more examples."},
}

func TestUnzip(t *testing.T) {

	setupZip(t)

	installer := PhpInstaller{
		"http://download.com",
		zipfile,
		temp,
	}

	err := installer.unzip()
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	for _, file := range files {
		path := temp + string(os.PathSeparator) + file.Name
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected %v; got %v", nil, err)
		}
	}

	tearDownZip(t)

}

func TestDownload(t *testing.T) {

	setupZip(t)

	fs := http.FileServer(http.Dir(output))
	http.Handle("/", fs)

	go http.ListenAndServe(":3031", nil)

	installer := PhpInstaller{
		"http://localhost:3031",
		zipfile,
		temp,
	}

	ioCom := IOCom{
		make(chan string),
		make(chan string),
		make(chan string),
		make(chan bool),
	}

	go func(ioCom IOCom) {
	Terminal:
		for {
			select {
			case outmsg := <-ioCom.Stdout:
				os.Stdout, _ = os.Open(os.DevNull)
				fmt.Print(outmsg)
			case <-ioCom.Done:
				break Terminal
			}
		}
	}(ioCom)

	err := installer.download(ioCom)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}
	ioCom.Done <- true

	if _, err := os.Stat(output); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}
}

func setupZip(t *testing.T) {

	os.MkdirAll(temp, 0755)
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			t.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		t.Fatal(err)
	}

	ioutil.WriteFile(output, buf.Bytes(), 0755)
}

func tearDownZip(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(temp, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(temp)
}
