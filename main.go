package main

import (
	"flag"

	"github.com/marcomilon/ezphp/app"
	"github.com/marcomilon/ezphp/engine/php"
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

	arguments := php.Arguments{
		*host,
		*docRoot,
	}

	phpInstaller := php.NewPhpInstaller()
	phpServer := php.NewPhpServer(arguments)

	go app.StartTerminal(ioChannels)
	app.Start(phpInstaller, phpServer, ioChannels, arguments)

}
