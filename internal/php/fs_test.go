package php_test

import (
	"os"
	"testing"

	"github.com/marcomilon/ezphp/internal/php"
)

var (
	phpNotFound = os.TempDir() + "nofound.exe"
	fakePHP     = os.TempDir() + "php.txt"
)

func TestFindExec(t *testing.T) {

	path := "echo"
	pathToEcho := "/bin/echo"

	absolutepath, err := php.FindExec(path)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if absolutepath != pathToEcho {
		t.Errorf("expected %v; got %v", pathToEcho, absolutepath)
	}

}

func TestFindPHP(t *testing.T) {

	setupFsTest(t)

	srv := php.Server{
		Exec:    phpNotFound,
		Host:    ":3001",
		DocRoot: "/var/www",
	}

	_, err := php.FindPHP(srv)
	if err == nil {
		t.Errorf("expected %v; got %v", err, nil)
	}

	srv2 := php.Server{
		Exec:    fakePHP,
		Host:    ":3001",
		DocRoot: "/var/www",
	}

	absolutepath, err := php.FindPHP(srv2)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if absolutepath != fakePHP {
		t.Errorf("expected %v; got %v", fakePHP, absolutepath)
	}

	teardownFsTest(t)

}

func setupFsTest(t *testing.T) {
	file, err := os.OpenFile(fakePHP, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
}

func teardownFsTest(t *testing.T) {
	err := os.Remove(fakePHP)
	if err != nil {
		t.Fatal(err)
	}
}
