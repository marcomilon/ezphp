package php

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {

	osSeparator := string(os.PathSeparator)
	pathZip := "unzip" + string(os.PathSeparator) + "ziptest"
	pathText1 := "unzip" + osSeparator + "ziptest" + osSeparator + "text1.txt"
	pathText2 := "unzip" + osSeparator + "ziptest" + osSeparator + "text2.txt"
	pathText3 := "unzip" + osSeparator + "ziptest" + osSeparator + "folder" + osSeparator + "text3.txt"

	if err := clearDir(pathZip); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	installer := PhpInstaller{
		"http://download.com",
		"ziptest.zip",
		"unzip",
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
