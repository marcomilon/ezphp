import QtQuick 2.1
import QtQuick.Controls 2.5

ApplicationWindow {
	id: window
	
	visible: true
	title: "EzPHP"
	width: 800
	height: 600
	
	ScrollView {
		id: view
		anchors.fill: parent
		
		TextArea {
			text: "TextArea"
			//readOnly: true
			color: "white"
			selectByMouse: true
			background: Rectangle { 
				anchors.fill: parent
				color: "#282c34" 
			}
		}
	}
	
}