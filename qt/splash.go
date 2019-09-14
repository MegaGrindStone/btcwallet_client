package qt

import (
	"log"

	"github.com/MegaGrindStone/btcd/rpcclient"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

// SplashCtxObj Structure
type SplashCtxObj struct {
	core.QObject
	_ func() `constructor:"init"`
	_ string `property:"status"`
}

func (s *SplashCtxObj) init() {
	s.SetStatus("Getting Headers")
}

// Load splash
func (s *SplashCtxObj) Load(c *rpcclient.Client) {
	var app = qml.NewQQmlApplicationEngine(nil)
	app.RootContext().SetContextProperty("ctxObject", s)
	app.Load(core.NewQUrl3("qt/qml/splash.qml", 0))

	go s.getHeaders(c)
}

func (s *SplashCtxObj) getHeaders(c *rpcclient.Client) {
	s.SetStatus("Mining")
	hashes, err := c.Generate(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, hash := range hashes {
		s.SetStatus(hash.String())
	}
	// accounts, err := c.ListAccounts()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// str := ""
	// for account, amount := range accounts {
	// 	str += " | " + fmt.Sprintf("%s : %s", account, amount.String())
	// }
	// s.SetStatus(str)
	// address, err := c.GetNewAddress("default")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s.SetStatus(address.EncodeAddress())
	// fmt.Println(address.EncodeAddress())
	// fmt.Println(address.String())
	s.SetStatus("Done")
}
