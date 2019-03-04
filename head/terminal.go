package head

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marcomilon/ezphp/engine"
)

func StartTerminal(ioCom engine.IOCom) {
Terminal:
	for {
		select {
		case outmsg := <-ioCom.Outmsg:
			fmt.Print(outmsg)
		case errMsg := <-ioCom.Errmsg:
			fmt.Print(errMsg)
		case <-ioCom.Done:
			byebye()
			break Terminal
		}
	}

	os.Exit(0)
}

func byebye() {
	fmt.Println("\nPress 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Confirm(question string) bool {

	var confirmation string

	fmt.Printf("%s [y/N]? ", question)
	fmt.Scanln(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		return true
	}

	return false
}
