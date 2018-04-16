package gtkui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type Ui struct {
    Tv *gtk.TextView
}

func (ui Ui) Write(b []byte) (n int, err error) {    
    s := string(b[0:])
    buffer, err := ui.Tv.GetBuffer()
    buffer.InsertAtCursor(s)
    return len(b), err
}

func (ui Ui) Show() {
    
	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile("gui.glade")

	obj, err := builder.GetObject("mainWindow")
	if err != nil {
		log.Fatal("Unable to find window:", err)
	}

	tvObj, err := builder.GetObject("logview")
	if err != nil {
		log.Fatal("Unable to find textview:", err)
	}

	tv := tvObj.(*gtk.TextView)
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}

	buffer.SetText("[EzPhp] Launching to EzPHP\n")
	buffer.InsertAtCursor("[About] https://github.com/marcomilon/ezphp\n")
    ui.Tv = tv

	if win, okwin := obj.(*gtk.Window); okwin {

		win.Connect("destroy", func() {
			gtk.MainQuit()
		})
		win.SetTitle("EzPHP")
		win.SetDefaultSize(800, 600)
		win.ShowAll()
        
	} else {
        
		log.Fatal("Unable to create window:", err)
	}

	gtk.Main()
}
