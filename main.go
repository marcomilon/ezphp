package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/marcomilon/ezphp/cli"
	"github.com/marcomilon/ezphp/engine/ezargs"
)

var DEBUG = "YES"

func main() {

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")
	installDir := flag.String("installDir", "localPHP", "<directory> - Installation directory for PHP.")
	about := flag.Bool("about", false, "Display about information.")

	flag.Parse()

	ezargs := ezargs.Arguments{
		*host,
		*docRoot,
		*installDir,
		*about,
	}

	if DEBUG == "YES" {

		file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Println("Unable to start EzPHP")
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)

	} else {

		log.SetOutput(ioutil.Discard)

	}

	cli.Start(ezargs)

}
