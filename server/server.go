package server

import (
    "fmt"
    "os/exec"
    "os"
)

type Serve struct {
    Path string
    Host string
    Port string
    DocumentRoot string
    cmd *exec.Cmd
}

func (s Serve) Start()  {    
    s.cmd = exec.Command(s.Path, "-S", s.Host + ":" + s.Port, "-t", s.DocumentRoot)
    s.cmd.Stdout = os.Stdout
    s.cmd.Stderr = os.Stderr
    execErr := s.cmd.Start()
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