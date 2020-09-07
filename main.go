package main

import (
	"flag"

	"github.com/marcomilon/ezphp/clients"
	"github.com/marcomilon/ezphp/ezphp"
	"github.com/marcomilon/ezphp/internal/php"
)

const (
	ezPHPVersion = "1.2.0"
	ezPHPWebsite = "https://github.com/marcomilon/ezphp"
)

func main() {

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")

	flag.Parse()

	ioChannels := php.IOCom{
		make(chan string),
		make(chan string),
		make(chan string),
		make(chan bool),
	}

	go clients.StartTerminal(ioChannels)

	instl := php.Installer{
		Source:      "https://windows.php.net/downloads/releases/php-7.4.10-nts-Win32-vc15-x64.zip",
		Destination: "php",
	}

	srv := php.Server{
		Exec:    "php/php.exe",
		Host:    *host,
		DocRoot: *docRoot,
	}

	ioChannels.Stdout <- "EzPHP v" + ezPHPVersion + "\n"
	ioChannels.Stdout <- ezPHPWebsite + "\n"
	ioChannels.Stdout <- "\n"

	ezphp.Start(srv, instl, ioChannels)

}
