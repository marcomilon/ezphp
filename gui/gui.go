package gui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/marcomilon/ezphp/serve"
	"log"
)

func Show() {

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

	tv, _ := tvObj.(*gtk.TextView)
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}

	buffer.SetText("[EzPhp] Launching to EzPHP\n")
	buffer.InsertAtCursor("[About] https://github.com/marcomilon/ezphp\n")

	if win, okwin := obj.(*gtk.Window); okwin {

		win.Connect("destroy", func() {
			gtk.MainQuit()
		})
		win.SetTitle("EzPHP")
		win.SetDefaultSize(800, 600)
		win.ShowAll()

		out := make(chan string)
		go serve.Serve(out)
		go func() {
			for {
				s := <-out
				buffer.InsertAtCursor(s)
			}
		}()

	} else {
        
		log.Fatal("Unable to create window:", err)
	}

	gtk.Main()
}
