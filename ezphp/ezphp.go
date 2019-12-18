package ezphp

import (
	"fmt"

	"github.com/marcomilon/ezphp/internal/php"
)

func Start(srv php.Server, instl php.Installer, ioCom php.IOCom) {

	_, err := php.FindPHP(srv)
	if err != nil {
		ioCom.Stdin <- "Would you like to install PHP?"
		result := <-ioCom.Stdin

		if result == "No" {
			ioCom.Done <- true
		}

		err := instl.Install(ioCom)
		if err != nil {
			ioCom.Stdout <- "Unable to install PHP\n"
			ioCom.Done <- true
		}

		ioCom.Stdout <- "PHP Installed succefully"
	}

	fmt.Printf(srv.DocRoot)
	if !php.Exists(srv.DocRoot) {
		if err := php.Create(srv.DocRoot); err != nil {
			ioCom.Stdout <- fmt.Sprintf("Unable to create document root, got: %v\n", err.Error())
			ioCom.Done <- true
		}

	}

}
