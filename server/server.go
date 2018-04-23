package server

import (
	"fmt"
	"io"
	"os/exec"
)

type Args struct {
	Php    string
	Host   string
	Public string
}

func Run(args Args, stdout io.Writer, stderr io.Writer) (*exec.Cmd, error) {
	fmt.Println("[Info] " + args.Php + " -S " + args.Host + " -t " + args.Public)
	cmd := exec.Command(args.Php, "-S", args.Host, "-t", args.Public)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(stdout, "[Error] Unable to execute PHP: " + err.Error())
		fmt.Fprintln(stdout, "[Info] php require to have the Visual C++ Redistributable for Visual Studio 2017")
		fmt.Fprintln(stdout, "[Info] Download Visual C++ from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145")
        return nil, err
	}
	return cmd, nil
}
