package php

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {

	osSeparator := string(os.PathSeparator)
	pathZip := "test" + osSeparator + "unzip" + osSeparator + "ziptest"
	pathText1 := pathZip + osSeparator + "text1.txt"
	pathText2 := pathZip + osSeparator + "text2.txt"
	pathText3 := pathZip + osSeparator + "folder" + osSeparator + "text3.txt"

	if err := clearDir(pathZip); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	installer := PhpInstaller{
		"http://download.com",
		"ziptest.zip",
		"test/unzip",
	}

	err := installer.unzip()
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if _, err := os.Stat(pathText1); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if _, err := os.Stat(pathText2); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if _, err := os.Stat(pathText3); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestDownload(t *testing.T) {

	osSeparator := string(os.PathSeparator)
	pathDownload := "test" + osSeparator + "download" + osSeparator + "ziptest.zip"

	os.Remove(pathDownload)

	fs := http.FileServer(http.Dir("test/unzip"))
	http.Handle("/", fs)

	go http.ListenAndServe(":3031", nil)

	installer := PhpInstaller{
		"http://localhost:3031",
		"ziptest.zip",
		"test/download",
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

	if _, err := os.Stat(pathDownload); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}
}

func clearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}
