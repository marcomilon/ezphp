package main

import (
    "flag"
    "github.com/marcomilon/ezphp/server"
    "github.com/marcomilon/ezphp/gtkui"
    "os"
)

type writer struct {
    msg chan string
}

func (w writer) Write(b []byte) (n int, err error) {
	s := string(b[0:])
    w.msg <- s
	return len(b), err
}

func main() {
    
    gui := flag.Int("gui", 0, "Path to php executable")
    path := flag.String("path", "", "Path to php executable")
    host := flag.String("host", "localhost", "Listening address")
    port := flag.String("port", "8080", "Listening port")
    documentRoot := flag.String("documentRoot", "web", "Specify document root")
    
    flag.Parse()
    
    s := server.Serve{}
    s.Path = *path
    s.Host = *host
    s.Port = *port
    s.DocumentRoot = *documentRoot
        
    if *gui == 1 {
        msg := make(chan string)
        w := writer{msg: msg}
        go gtkui.Show(msg)
        s.Stdout = w
        s.Stderr = w
        s.Run()
    } else {
        s.Stdout = os.Stdout
        s.Stderr = os.Stderr
        s.Run()
    }
    
    // c := 2;
    // 
    // if c > 5 {
    //     cli.Start()
    // } else {       
    //     gui.Start()
    // }
}
