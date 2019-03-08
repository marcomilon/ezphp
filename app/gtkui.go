package app

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/marcomilon/ezphp/engine/php"
)

func StartWin() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("EzPHP")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("EzPHP")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	vbox.PackStart(l, true, false, 0)

	win.Add(vbox)
	win.SetDefaultSize(800, 600)
	win.ShowAll()

	dialog := gtk.MessageDialogNew(win, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_QUESTION, gtk.BUTTONS_OK_CANCEL, "Would you like to install PHP?")
	dialog.SetTitle("EzPHP")
	dialog.Run()

	gtk.Main()
}

func StartUI(ioCom php.IOCom) {
	gtk.Init(nil)

	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("EzPHP")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	css, err := gtk.CssProviderNew()
	if nil != err {
		log.Fatal("Unable to create css provider: ", err)
	}

	err2 := css.LoadFromPath("./ezphp.css")
	if nil != err2 {
		log.Fatal("Unable to load css: ", err2)
	}

	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	sw, err := gtk.ScrolledWindowNew(nil, nil)
	if nil != err {
		log.Fatal("Unable to create label: ", err)
	}

	tv, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create label: ", err)
	}

	go func(ioCom php.IOCom, tv *gtk.TextView) {
	Gui:
		for {
			select {
			case outmsg := <-ioCom.Outmsg:
				buffer, _ := tv.GetBuffer()
				buffer.InsertAtCursor(outmsg)
			case errMsg := <-ioCom.Errmsg:
				buffer, _ := tv.GetBuffer()
				buffer.InsertAtCursor(errMsg)
			case <-ioCom.Done:
				break Gui
			}
		}
	}(ioCom, tv)

	sw.Add(tv)
	win.Add(sw)
	win.SetDefaultSize(800, 600)
	win.ShowAll()

	gtk.Main()
}
