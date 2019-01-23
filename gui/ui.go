package gui

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
)

func StartUI() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML
	gui.NewQGuiApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Fusion")

	// // create the qml application engine
	// engine := qml.NewQQmlApplicationEngine(nil)
	//
	// // load the embedded qml file
	// // created by either qtrcc or qtdeploy
	// // engine.Load(core.NewQUrl3("qrc:/qml/ui.qml", 0))
	// // you can also load a local file like this instead:
	// engine.Load(core.QUrl_FromLocalFile("./gui/qml/ui.qml"))

	view := quick.NewQQuickView(nil)
	view.SetMinimumSize(core.NewQSize2(800, 600))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("EzPHP")

	view.SetSource(core.QUrl_FromLocalFile("./gui/qml/ui.qml"))

	input := view.RootObject().FindChild("outText", core.Qt__FindChildrenRecursively)
	input.SetProperty("text", core.NewQVariant14("EzPHP\n"))

	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gui.QGuiApplication_Exec()
}
