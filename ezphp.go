package main

import (
	"flag"
	"github.com/marcomilon/ezphp/gtkui"
	"github.com/marcomilon/ezphp/server"
	"os"
	"os/exec"
)

func main() {

	ui := flag.String("ui", "cli", "Launch ezphp with gui or console <gui:cli>")
	php := flag.String("php", "/usr/bin/php", "Path to php executable")
	host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
	public := flag.String("public", "web", "Path to public directory")

	flag.Parse()

	args := server.Args{
		Php:    *php,
		Host:   *host,
		Public: *public,
	}

	uiInterface := *ui

	switch uiInterface {
	case "gui":

		var cmd *exec.Cmd
        var err error
        
		finished := make(chan bool)

		gui := gtkui.NewGui()

		go func() {
			gui.Show(finished)
		}()

		go func() {
			cmd, err = server.Run(args, gui, gui)
		}()
        
		<-finished
        if(err == nil) {
		    cmd.Process.Kill()
        }

	default:
		server.Run(args, os.Stdout, os.Stdin)
	}
    
    return
}
