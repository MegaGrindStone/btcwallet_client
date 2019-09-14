import QtQuick 2.9
import QtQuick.Window 2.2

Window {
    id: window
    visible: true
    width: 400
    height: 400
    flags: Qt.FramelessWindowHint

    Text {
        id: element
        x: 187
        y: 141
        width: 267
        height: 130
        text: qsTr("Wallet")
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
        font.family: "Verdana"
        lineHeight: 0
        fontSizeMode: Text.FixedSize
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        clip: false
        font.pixelSize: 72
    }

    Text {
        id: element1
        x: 240
        y: 423
        text: ctxObject.status
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 10
        font.pixelSize: 14
    }
}
