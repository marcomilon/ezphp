package output

import (
    "fmt"
    "log"
	"github.com/gotk3/gotk3/gtk"
)

type Output struct {
	Tv *gtk.TextView
}

func (output Output) Write(b []byte) (n int, err error) {
	//s := string(b[:n])
	fmt.Sprintf(">> %s", b)
    buffer, err := output.Tv.GetBuffer()
    if err != nil {
        log.Fatal("Unable to get buffer:", err)
    }
    
    buffer.InsertAtCursor("hello\n")
	return len(b), err
}
