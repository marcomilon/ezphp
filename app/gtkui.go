package app

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/marcomilon/ezphp/engine/php"
)

var listenToKeyEvents = true
var freezeTextView = false
var TextViewText = ""

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

	tv.Connect("key-press-event", func(tv *gtk.TextView, ev *gdk.Event) {

		if listenToKeyEvents {

			listenToKeyEvents = false
			freezeTextView = true

			keyEvent := &gdk.EventKey{ev}

			if keyEvent.KeyVal() == gdk.KEY_y {
				ioCom.Confirm <- "Yes"
			} else {
				ioCom.Confirm <- "No"
			}

		}

	})

	go func(ioCom php.IOCom, tv *gtk.TextView) {
	Gui:
		for {
			select {
			case outmsg := <-ioCom.Outmsg:
				sendToTv(tv, outmsg)
			case errMsg := <-ioCom.Outmsg:
				sendToTv(tv, errMsg)
			case confirmMsg := <-ioCom.Confirm:
				input := php.NewStdin(confirmMsg + " [y/N] ")
				sendToTv(tv, input)
			case <-ioCom.Done:
				outmsg := php.NewStdout("\n\nClose window to exit...")
				sendToTv(tv, outmsg)
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

func sendToTv(tv *gtk.TextView, outmsg php.IOMessage) {
	glib.IdleAdd(func() {
		buffer, _ := tv.GetBuffer()

		if outmsg.IOContext == php.STDINSTALL {
			freezeTextView = true
		} else {
			freezeTextView = false
		}

		if freezeTextView {

			if TextViewText == "" {
				buffer.InsertAtCursor(outmsg.Msg)
				start, end := buffer.GetBounds()
				TextViewText, _ = buffer.GetText(start, end, true)
			} else {
				buffer.SetText(TextViewText + outmsg.Msg)
			}

		} else {
			buffer.InsertAtCursor(outmsg.Msg)
		}
	})
}
