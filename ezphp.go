package main

import (
    "flag"
    "github.com/marcomilon/ezphp/server"
    "github.com/marcomilon/ezphp/gtkui"
    "os"
)

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
        ui := gtkui.Ui{}
        go ui.Show()
        s.Stdout = ui
        s.Stderr = ui
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
