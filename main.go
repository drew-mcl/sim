package main

import (
	"log"
	"time"

	"github.com/drew-mcl/sim/pkg/metrics"
	"github.com/drew-mcl/sim/pkg/simulator"
)

func main() {
	// Initialize components
	histogram := metrics.NewLatencyHistogram(5 * time.Minute)
	receiver := simulator.NewReceiver(histogram)
	sender := simulator.NewSender(receiver)

	// Start simulation
	sender.StartSending()
	defer sender.StopSending()

	// Output histogram periodically
	for range time.Tick(1 * time.Minute) {
		log.Println("Current Histogram:", histogram.Report())
	}
}
