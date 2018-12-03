package main

import (
	"flag"
	"log"
	"os"

	"github.com/marcomilon/ezphp/cli"
	"github.com/marcomilon/ezphp/engine/ezargs"
)

func main() {

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")
	installDir := flag.String("installDir", "localPHP", "<directory> - Installation directory for PHP.")

	flag.Parse()

	ezargs := ezargs.Arguments{
		*host,
		*docRoot,
		*installDir,
	}

	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("check to make sure it works")

	cli.Start(ezargs)

	return
}
