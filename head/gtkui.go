package head

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type GtkUI struct {
	Tv *gtk.TextView
}

func (guiIO GtkUI) Write(b []byte) (int, error) {
	s := string(b[0:])

	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)

	return len(s), nil
}

func (guiIO GtkUI) Info(s string) {
	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)
}

func (guiIO GtkUI) Error(s string) {
	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)
}

func (guiIO GtkUI) Custom(tag string, s string) {
	fmt.Print(s)
}

func (guiIO GtkUI) Confirm(question string) bool {

	return false
}

func StartUI() {
	gtk.Main()
}

func SetupUI() GtkUI {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("EzPHP")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	sw, err := gtk.ScrolledWindowNew(nil, nil)
	if nil != err {
		log.Fatal("Unable to create label:", err)
	}

	tv, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	sw.Add(tv)
	win.Add(sw)
	win.SetDefaultSize(800, 600)
	win.ShowAll()

	return GtkUI{Tv: tv}
}
