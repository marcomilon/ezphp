package main

import (
	"flag"
	"log"
	"os"

	"github.com/marcomilon/ezphp/app"
	"github.com/marcomilon/ezphp/engine/ezargs"
	"github.com/marcomilon/ezphp/engine/php"
	"github.com/sirupsen/logrus"
)

var DEBUG = "YES"

func init() {

	file, err := os.OpenFile("ezphp.log", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("Unable to start EzPHP")
		log.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetOutput(file)
}

func main() {

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")

	flag.Parse()

	ezargs := ezargs.Arguments{
		*host,
		*docRoot,
	}

	ioChannels := php.IOCom{
		Outmsg:  make(chan php.IOMessage),
		Confirm: make(chan string),
		Done:    make(chan bool),
	}

	go app.StartTerminal(ioChannels)
	app.Start(ezargs, ioChannels)

}
