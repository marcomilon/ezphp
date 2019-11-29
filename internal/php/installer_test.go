package php_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/marcomilon/ezphp/internal/php"
)

var (
	temp        = os.TempDir() + "ezphptest"
	zipfile     = temp + string(os.PathSeparator) + "out.zip"
	source      = "http://localhost:3031/out.zip"
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

	setupZip(t)
	ioCom := setup()

	err := ioCom.Install(source, destination)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	for _, file := range files {
		path := temp + string(os.PathSeparator) + file.name
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected %v; got %v", nil, err)
		}
	}

	tearDownZip(t)
}

func setup() php.IOCom {
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
				os.Stdout, _ = os.Open(os.DevNull)
				fmt.Print(outmsg)
			case <-ioCom.Done:
				break Terminal
			}
		}
	}(ioCom)

	return ioCom
}

func setupZip(t *testing.T) {

	os.MkdirAll(temp, 0755)
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	for _, file := range files {
		f, err := w.Create(file.name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write([]byte(file.body))
		if err != nil {
			t.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		t.Fatal(err)
	}

	ioutil.WriteFile(zipfile, buf.Bytes(), 0755)
}

func tearDownZip(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(temp, "*"))

	if err != nil {
		t.Fatal(err.Error())
	}

	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	os.Remove(temp)
}
