package docroot_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/marcomilon/ezphp/engine/fs/docroot"
)

var path string = os.TempDir() + "ezphptest"

func TestCreate(t *testing.T) {

	setup(t)

	err := docroot.Create(path)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}
}

func TestExists(t *testing.T) {

	setup(t)

	var exists bool

	exists = docroot.Exists(path)
	if exists {
		t.Errorf("expected %v; got %v", false, exists)
	}

	os.MkdirAll(path, 0755)
	exists = docroot.Exists(path)
	if !exists {
		t.Errorf("expected %v; got %v", true, exists)
	}

}

func TestCreateIndex(t *testing.T) {

	setup(t)

	os.MkdirAll(path, 0755)
	osSeparator := string(os.PathSeparator)
	pathIndex := path + osSeparator + "index.php"

	template := `index`

	err := docroot.CreateIndex(pathIndex, template)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if _, err := os.Stat(pathIndex); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index, err := ioutil.ReadFile(pathIndex)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if template != string(index) {
		t.Errorf("expected %v; got %s", template, index)
	}

	tearDown(t)
}

func setup(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(path)
}

func tearDown(t *testing.T) {
	setup(t)
}
