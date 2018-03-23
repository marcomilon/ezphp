package main

import (
    "fmt"
    "os/exec"
    "github.com/marcomilon/ezphp/installer"
)

func main() {
    
    fmt.Printf("EzPHP\n")
    
    path, err := searchPhpBin()
    if err != nil {
        fmt.Printf("php not found\n")
        path, err = installer.Install()
    } 
    
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("php is available at %s\n", path);
    
    // args := []string{"php", "-S", "localhost:8888"}
    // env := os.Environ()
    // 
    // execErr := syscall.Exec(path, args, env)
    // if execErr != nil {
    //     panic(execErr)
    // }
}

func searchPhpBin() (string, error) {
    path, err := exec.LookPath(installer.PhpExecutable)
    if err != nil {
        return "", err 
    }
    
    return path, nil
}