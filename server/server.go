package server

import (
	"fmt"
	"os/exec"
    "os"
)

type Args struct {
	Php    string
	Host   string
	Public string
}

func Run(args Args) {
	fmt.Println("[Info] Running php server: " + args.Php + " -S " + args.Host + " -t " + args.Public)
	cmd := exec.Command(args.Php, "-S", args.Host, "-t", args.Public)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("[Error] Unable to execute PHP: " + err.Error())
		fmt.Println("[Info] php require to have the Visual C++ Redistributable for Visual Studio 2017")
		fmt.Println("[Info] Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145")
        fmt.Print("[Info] Press Enter to exit...")
        fmt.Scanln()
	}
}
