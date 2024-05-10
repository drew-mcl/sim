package simulator

import (
	"testing"
	"time"
)

func TestSender(t *testing.T) {
	histogram := metrics.NewLatencyHistogram(5 * time.Minute)
	receiver := NewReceiver(histogram)
	sender := NewSender(receiver)

	sender.StartSending()
	time.Sleep(3 * time.Second) // Allow some messages to be sent
	sender.StopSending()

	if len(histogram.Report()) == 0 {
		t.Error("No latencies recorded, expected at least one")
	}
}
