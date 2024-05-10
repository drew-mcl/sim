package metrics

import (
	"sync"
	"time"
)

type LatencyHistogram struct {
	windowDuration time.Duration
	data           []time.Duration
	lock           sync.Mutex
}

func NewLatencyHistogram(windowDuration time.Duration) *LatencyHistogram {
	return &LatencyHistogram{
		windowDuration: windowDuration,
		data:           []time.Duration{},
	}
}

func (h *LatencyHistogram) AddLatency(latency int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, time.Duration(latency)*time.Millisecond)
	h.cleanup()
}

func (h *LatencyHistogram) cleanup() {
	cutoff := time.Now().Add(-h.windowDuration)
	for i, latency := range h.data {
		if latency.Before(cutoff) {
			h.data = h.data[i+1:]
			break
		}
	}
}

func (h *LatencyHistogram) Report() []time.Duration {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.data
}
