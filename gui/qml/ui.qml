import QtQuick 2.1
import QtQuick.Controls 2.5


ScrollView {
	id: view
	anchors.fill: parent
	
	TextArea {
		objectName: "outText"
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