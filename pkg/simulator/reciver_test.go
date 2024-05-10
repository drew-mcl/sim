package simulator

import (
	"fix"
	"metrics"
	"testing"
	"time"
)

func TestReceiver(t *testing.T) {
	histogram := metrics.NewLatencyHistogram(5 * time.Minute)
	receiver := NewReceiver(histogram)

	message := fix.CreateMessage()
	receiver.Receive(message)

	if len(histogram.Report()) != 1 {
		t.Errorf("Expected 1 latency entry, got %d", len(histogram.Report()))
	}
}
