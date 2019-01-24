package gui

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func StartUI() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 600)
	window.SetWindowTitle("EzPHP")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// create a line edit
	// with a custom placeholder text
	// and add it to the central widgets layout
	input := widgets.NewQTextEdit(nil)
	input.SetStyleSheet("QTextEdit { background-color: black; color: white; font-size: 16px }")
	//input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	// make the window visible
	window.Show()

	input.Append("Ez")
	input.Append("Php")

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}
