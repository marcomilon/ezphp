package main

import (
    "flag"
    "time"
    "github.com/marcomilon/ezphp/server"
    "github.com/marcomilon/ezphp/gui"
)

func main() {
    
    useGui := flag.Int("gui", 1, "Path to php executable")
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
    
    
    if *useGui == 1 {
        gui.Start();
    }
    
    s.Start()
    
    time.Sleep(time.Millisecond * 500)
    
    s.Stop()
    
    // c := 2;
    // 
    // if c > 5 {
    //     cli.Start()
    // } else {       
    //     gui.Start()
    // }
}
