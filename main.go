package main

import (
	"flag"

	"github.com/marcomilon/ezphp/cli"
	"github.com/marcomilon/ezphp/engine/ezargs"
)

const (
	ezPHPVersion = "1.1.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
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

	cli.Clean(ezargs)
}
