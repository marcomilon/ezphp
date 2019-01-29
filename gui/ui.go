package gui

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type GuiIO struct {
	Tv *gtk.TextView
}

func (guiIO GuiIO) Write(b []byte) (int, error) {
	s := string(b[0:])

	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)

	return len(s), nil
}

func (guiIO GuiIO) Info(s string) {
	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)
}

func (guiIO GuiIO) Error(s string) {
	buffer, _ := guiIO.Tv.GetBuffer()
	buffer.InsertAtCursor(s)
}

func (guiIO GuiIO) Custom(tag string, s string) {
	fmt.Print(s)
}

func (guiIO GuiIO) Confirm(question string) bool {

	return false
}

func StartUI() {
	gtk.Main()
}

func SetupUI() GuiIO {
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

	return GuiIO{Tv: tv}
}
