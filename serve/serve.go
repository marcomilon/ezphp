package serve

import (
	"fmt"
	"github.com/marcomilon/ezphp/installer"
    "github.com/marcomilon/ezphp/output"
	"os/exec"
)

func Serve(output output.Output) {

	path, err := searchPhpBin()
	if err != nil {
		path, err = installer.PathToPhp()
		if err != nil {
			path, err = installer.Install()
		}
	}

	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}
    
    fmt.Printf("[EzPhp] Using php located in: %s\n", path)
    servePhp(path, output)

}

func servePhp(path string, o output.Output) {
	command := exec.Command(path, "-S", "localhost:" + installer.Port, "-t", installer.DocumentRoot)
	command.Stderr = o
	execErr := command.Start()
	if execErr != nil {
		fmt.Printf("[Error] Unable to execute PHP: %s\n", execErr.Error())
		fmt.Printf("[Error] php is located in: %s\n", path)
        fmt.Println("[Error] php require to have the Visual C++ Redistributable for Visual Studio 2017")
        fmt.Println("[Error] Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145")
	}
}

func searchPhpBin() (string, error) {
	path, err := exec.LookPath(installer.PhpExecutable)
	if err != nil {
		return "", err
	}

	return path, nil
}