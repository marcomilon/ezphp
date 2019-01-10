package gui

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type GuiIO struct {
	Tv *gtk.TextView
}

func (g GuiIO) Write(b []byte) (int, error) {
	s := string(b[0:])
	fmt.Print(s)

	return len(b), nil
}

func (g GuiIO) Info(s string) {
	buffer := get_buffer_from_tview(g.Tv)
	buffer.InsertAtCursor(s)
}

func (g GuiIO) Error(s string) {
	buffer := get_buffer_from_tview(g.Tv)
	buffer.InsertAtCursor(s)
}

func (g GuiIO) Custom(tag string, s string) {
	fmt.Print(s)
}

func (g GuiIO) Confirm(question string) bool {

	// var confirmation string
	//
	// io.Info(fmt.Sprintf("%s [y/N]? ", question))
	// fmt.Scanln(&confirmation)
	//
	// confirmation = strings.TrimSpace(confirmation)
	// confirmation = strings.ToLower(confirmation)
	//
	// if confirmation == "y" {
	// 	return true
	// }

	return false
}

func get_buffer_from_tview(tv *gtk.TextView) *gtk.TextBuffer {
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}
	return buffer
}
