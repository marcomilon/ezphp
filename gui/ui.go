package gui

import (
	"github.com/andlabs/ui"
)

func StartUI() {
	err := ui.Main(func() {
		window := ui.NewWindow("EzPHP", 800, 600, false)
		window.SetMargined(true)

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)

		label := ui.NewMultilineEntry()
        
        vbox.Append(label, true)

		window.SetChild(vbox)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
		//go counter(label)
	})
	if err != nil {
		panic(err)
	}
}

// func counter(label *ui.Label) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second)
// 		ui.QueueMain(func() {
// 			label.SetText("number " + strconv.Itoa(i))
// 		})
// 	}
// }
