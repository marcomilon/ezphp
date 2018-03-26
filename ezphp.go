package main

import (
	"fmt"
	"github.com/marcomilon/ezphp/serve"
    "github.com/gotk3/gotk3/gtk"
)

func main() {

	fmt.Println("[EzPhp] Launching to EzPHP")
	fmt.Println("[About] https://github.com/marcomilon/ezphp")

    serve.Serve()
}