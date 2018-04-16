package server

import (
    "fmt"
    "os/exec"
    "io"
)

type Serve struct {
    Path string
    Host string
    Port string
    DocumentRoot string
    Stdout io.Writer
    Stderr io.Writer
    cmd *exec.Cmd
}

func (s Serve) Run()  {    
    s.cmd = exec.Command(s.Path, "-S", s.Host + ":" + s.Port, "-t", s.DocumentRoot)
    s.cmd.Stdout = s.Stdout
    s.cmd.Stderr = s.Stderr
    execErr := s.cmd.Run()
    if execErr != nil {
        fmt.Println("[Error] Unable to execute PHP: %s", execErr.Error())
        fmt.Println("[Error] php is located in: %s", s.Path)
        fmt.Println("[Error] php require to have the Visual C++ Redistributable for Visual Studio 2017")
        fmt.Println("[Error] Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145")
    }
}

func (s Serve) Stop()  { 
    s.cmd.Process.Kill()
}