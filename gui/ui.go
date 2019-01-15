package gui

import (
	"fmt"

	"github.com/andlabs/ui"
)

type UiIO struct {
	TextArea *ui.MultilineEntry
}

func (uiIo UiIO) Write(b []byte) (int, error) {
	s := string(b[0:])
	fmt.Print(s)

	return len(b), nil
}

func (uiIo UiIO) Info(s string) {
	ui.QueueMain(func() {
		uiIo.TextArea.Append(s)
	})
}

func (uiIo UiIO) Error(s string) {
	ui.QueueMain(func() {
		uiIo.TextArea.Append(s)
	})
}

func (uiIo UiIO) Custom(tag string, s string) {
	fmt.Print(s)
}

func (uiIo UiIO) Confirm(question string) bool {

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

func StartUI() (*ui.Window, *ui.MultilineEntry) {
	err := ui.Main(func() {

		window := ui.NewWindow("EzPHP", 800, 600, true)
		vbox := ui.NewVerticalBox()
		label := ui.NewMultilineEntry()

		vbox.Append(label, true)

		window.SetChild(vbox)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
		
		
	})

	return label
}

// func counter(label *ui.Label) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second)
// 		ui.QueueMain(func() {
// 			label.SetText("number " + strconv.Itoa(i))
// 		})
// 	}
// }
