package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/marcomilon/ezphp/internal/php"
)

const (
	ezPHPWebsite  = "https://github.com/marcomilon/ezphp"
	downloadURL   = "https://windows.php.net/downloads/releases/php-7.4.10-nts-Win32-vc15-x64.zip"
	installFolder = "php-7.4.10"
)

func main() {

	var phpExe string
	var phpExeAbs string
	var err error

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")

	flag.Parse()

	fmt.Println("EzPHP (" + ezPHPWebsite + ")")

	phpExe, err = php.FindPHPExec(installFolder)
	if err != nil {
		phpExe, err = php.FastInstall(downloadURL, installFolder)

		if err != nil {
			fmt.Printf("Unable to install PHP: %v", err)
			php.ExitEzPHP()
		}

	}

	phpExeAbs, _ = filepath.Abs(phpExe)
	docRootAbs, _ := filepath.Abs(*docRoot)

	if _, err := os.Stat(*docRoot); os.IsNotExist(err) {
		os.Mkdir(*docRoot, 0755)
	}

	fmt.Println("")
	fmt.Println("** Information **")
	fmt.Println("Copy your PHP files in the Document root folder.")
	fmt.Println("")
	fmt.Printf("%-16s | %s\n", "PHP located in", phpExeAbs)
	fmt.Printf("%-16s | %s\n", "Document root", docRootAbs)
	fmt.Printf("%-16s | %s\n", "Website URL", fmt.Sprintf("%s%s", "http://", *host))
	fmt.Println("")

	cmd := exec.Command(phpExe, "-S", *host, "-t", *docRoot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		log.Fatal(err)
	}

}
