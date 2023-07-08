package rates

import (
	"fmt"
	"log"

	"github.com/hadrianl/ibapi"
)

type StreamRatesApp struct {
	Client *ibapi.IbClient
}

func NewStreamRatesApp() *StreamRatesApp {
	ibwrapper := &ibapi.Wrapper{}
	app := &StreamRatesApp{
		Client: ibapi.NewIbClient(ibwrapper),
	}
	return app
}

func (app *StreamRatesApp) Connect() {
	err := app.Client.Connect("127.0.0.1", 7497, 0)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
}

func (app *StreamRatesApp) Disconnect() {
	app.Client.Disconnect()
}

func (app *StreamRatesApp) OnTickPrice(tick *ibapi.TickPrice) {
	if tick.TickType == 2 { // Bid price
		fmt.Printf("Bid Price: %v\n", tick.Price)
	} else if tick.TickType == 4 { // Ask price
		fmt.Printf("Ask Price: %v\n", tick.Price)
	}
}
