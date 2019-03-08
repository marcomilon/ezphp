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
	installDir := flag.String("installDir", "localPHP", "<directory> - Installation directory for PHP.")
	useGui := flag.Bool("gui", false, "Use gui interface.")

	flag.Parse()

	ezargs := ezargs.Arguments{
		*host,
		*docRoot,
		*installDir,
		*useGui,
	}

	ioChannels := php.IOCom{
		Outmsg:  make(chan string),
		Errmsg:  make(chan string),
		Confirm: make(chan string),
		Done:    make(chan bool),
	}

	if ezargs.Gui {
		//go app.StartUI(ioChannels)
		app.StartWin()
		//app.Start(ezargs, ioChannels)
	} else {
		go app.StartTerminal(ioChannels)
		app.Start(ezargs, ioChannels)
	}

}
