package main

import (
	"flag"
	"os"

	"github.com/marcomilon/ezphp/internals/install"
	"github.com/marcomilon/ezphp/internals/output"
)

var (
	path string
	err  error
)

func main() {

	destination := flag.String("destination", "", "Folder to install php")
	version := flag.String("version", "", "Php version to install")

	flag.Parse()

	if *version == "" || *destination == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	output.Installer("Installing " + *version + " Please wait...\n")

	path, err = install.Installer(*version, *destination)

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.Info("Php was installed in directiory: " + path)
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
