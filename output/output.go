package output

import (
    "log"
	"github.com/gotk3/gotk3/gtk"
)

type Output struct {
	Tv *gtk.TextView
}

func (output Output) Write(b []byte) (n int, err error) {
	s := string(b[1:])
    
    buffer, err := output.Tv.GetBuffer()
    if err != nil {
        log.Fatal("Unable to get buffer:", err)
    }
    
    buffer.InsertAtCursor(s)
	return len(b), err
}
