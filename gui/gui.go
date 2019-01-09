package gui

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	Builder *gtk.Builder
	Tv      *gtk.TextView
	Window  *gtk.Window
}

func (g *Gui) Write(b []byte) (n int, err error) {
	s := string(b[0:])
	buffer, err := g.Tv.GetBuffer()
	buffer.InsertAtCursor(s)
	return len(b), err
}

func (g *Gui) Show() {

	g.Window.Connect("destroy", func() {
		fmt.Println("Bye Bye")
	})
	g.Window.SetTitle("EzPHP")
	g.Window.SetDefaultSize(800, 600)
	g.Window.ShowAll()

	gtk.Main()
}

func NewGui() *Gui {

	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile("gui.glade")
	if err != nil {
		log.Fatal("Unable to load gui.glade:", err)
	}

	winObj, err := builder.GetObject("mainWindow")
	if err != nil {
		log.Fatal("Unable to find window:", err)
	}

	tvObj, err := builder.GetObject("logview")
	if err != nil {
		log.Fatal("Unable to find textview:", err)
	}

	tv := tvObj.(*gtk.TextView)
	win := winObj.(*gtk.Window)

	return &Gui{
		Builder: builder,
		Tv:      tv,
		Window:  win,
	}

}

func Start() {
	logrus.Debug("Starting GUI")
	gui := NewGui()
	gui.Show()
}
