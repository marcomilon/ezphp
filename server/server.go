package server

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/marcomilon/ezphp/internals/output"
)

type Args struct {
	Php    string
	Host   string
	Public string
}

func Run(args Args) {

	output.Info("Command to launch php server: " + args.Php + " -S " + args.Host + " -t " + args.Public + "\n")
	output.Info("Your server url is: " + "http://" + args.Host + "\n")
	cmd := exec.Command(args.Php, "-S", args.Host, "-t", args.Public)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		output.Error("Unable to execute PHP: " + err.Error() + "\n")
		output.Info("php require to have the Visual C++ Redistributable for Visual Studio 2017\n")
		output.Info("Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145\n")
		output.Info("Press Enter to exit... ")
		fmt.Scanln()
	}
}
