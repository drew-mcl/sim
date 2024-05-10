package simulator

import (
	"fix"
	"time"
)

type Sender struct {
	receiver *Receiver
	stopCh   chan bool
}

func NewSender(receiver *Receiver) *Sender {
	return &Sender{
		receiver: receiver,
		stopCh:   make(chan bool),
	}
}

func (s *Sender) StartSending() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				message := fix.CreateMessage()
				s.receiver.Receive(message)
			case <-s.stopCh:
				ticker.Stop()
				return
			}
		}
	}()
}

func (s *Sender) StopSending() {
	close(s.stopCh)
}
