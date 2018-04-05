package serve

import (
	"fmt"
	"github.com/marcomilon/ezphp/installer"
	"os/exec"
)

type writer struct {
    outChan chan<- string
}

func (w writer) Write(b []byte) (n int, err error) {
	s := string(b[0:])
    w.outChan <- s
	return len(b), err
}

func Serve(out chan<- string) {
    w := writer{outChan: out}
    path, _ := searchPhpBin()
    
	command := exec.Command(path, "-S", "localhost:" + installer.Port, "-t", installer.DocumentRoot)
    command.Stdout = w
	command.Stderr = w
    execErr := command.Start()
	if execErr != nil {
		fmt.Fprintln(w, "[Error] Unable to execute PHP: %s", execErr.Error())
		fmt.Fprintln(w, "[Error] php is located in: %s", path)
        fmt.Fprintln(w, "[Error] php require to have the Visual C++ Redistributable for Visual Studio 2017")
        fmt.Fprintln(w, "[Error] Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145")
	}
}

func searchPhpBin() (string, error) {
	path, err := exec.LookPath(installer.PhpExecutable)
	if err != nil {
		return "", err
	}

	return path, nil
}