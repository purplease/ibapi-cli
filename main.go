package main

import (
	"ibapi-cli/pkg/rates"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hadrianl/ibapi"
)

func main() {
	// Create a channel to receive the SIGTERM signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	app := rates.NewStreamRatesApp()
	app.Connect()

	// Subscribe to tick price updates
	app.Client.OnTickPrice(app.onTickPrice)

	// Create the contract for EUR/USD CFD
	contract := ibapi.Contract{
		Symbol:   "EUR",
		SecType:  "CFD",
		Currency: "USD",
		Exchange: "SMART",
	}

	// Request streaming data for the contract
	app.Client.RequestMarketData(1, &contract, "", false, nil)

	// Start a goroutine to wait for the SIGTERM signal
	go func() {
		<-sigChan
		log.Println("Received SIGTERM signal. Closing the application...")
		app.Disconnect()
		os.Exit(0)
	}()

	// Wait indefinitely
	select {}
}
