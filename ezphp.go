package main

import (
    "flag"
    "github.com/marcomilon/ezphp/server"
)

func main() {
    
    php := flag.String("php", "/usr/bin/php", "Path to php executable")
    host := flag.String("host", "localhost:8080", "Listening address: <addr>:<port> ")
    public := flag.String("public", "web", "Path to public directory")
    
    flag.Parse()
    
    args := server.Args{
        Php:    *php,
        Host:   *host,
        Public: *public,
    }
    
    server.Run(args)
    
    return
}
