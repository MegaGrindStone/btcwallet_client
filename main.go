package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	// connCfg := &rpcclient.ConnConfig{
	// 	Host:         "localhost:8332",
	// 	User:         "rpcuser",
	// 	Pass:         "rpcpass",
	// 	HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
	// 	DisableTLS:   true, // Bitcoin core does not provide TLS by default
	// }
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	// client, err := rpcclient.New(connCfg, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Shutdown()

	// Get the current block count.
	// blockCount, err := client.GetBlockCount()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Block count: %d", blockCount)

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qt/qml/splash.qml", 0))

	gui.QGuiApplication_Exec()
}
