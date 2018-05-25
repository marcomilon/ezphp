package main

import (
	"flag"
	"fmt"

	"github.com/marcomilon/ezphp/internals/output"
	"github.com/marcomilon/ezphp/internals/serve"
)

func main() {
	php := flag.String("php", "", "Path to php executable")
	host := flag.String("host", "", "Listening address: <addr>:<port> ")
	public := flag.String("public", "", "Path to public directory")

	flag.Parse()
	
	if *php == "" || *host == "" || *public == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	output.Info("Command: " + *php + " -S " + *host + " -t " + *public + "\n")
	output.Info("Your server url is: " + "http://" + *host + "\n")

	err := serve.Start(*php, *host, *public)

	if err != nil {
		output.Error("Unable to execute PHP: " + err.Error() + "\n")
		output.Error("php requires to have the Visual C++ Redistributable for Visual Studio 2017 installed\n")
		output.Error("If you don't have it download it from here: https://www.microsoft.com/en-us/download/details.aspx?id=48145\n")
		output.Error("Press Enter to continue... ")
		fmt.Scanln()
	}

}
