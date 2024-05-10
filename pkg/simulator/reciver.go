package simulator

import (
	"fix"
	"metrics"
	"time"
)

type Receiver struct {
	histogram *metrics.LatencyHistogram
}

func NewReceiver(histogram *metrics.LatencyHistogram) *Receiver {
	return &Receiver{histogram: histogram}
}

func (r *Receiver) Receive(message fix.Message) {
	start := time.Now()
	// Simulate processing time
	time.Sleep(100 * time.Millisecond)
	latency := time.Since(start)
	r.histogram.AddLatency(int(latency.Milliseconds()))
}
