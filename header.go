package swaybarprotocol

import (
	"syscall"
)

// Header represents the swaybar header as defined in the swaybar
// json protocol. (see: man 7 swaybar-protocol)
type Header struct {
	Version     int            `json:"version"`
	ClickEvents bool           `json:"click_events,omitempty"`
	ContSignal  syscall.Signal `json:"cont_signal,omitempty"`
	StopSignal  syscall.Signal `json:"stop_signal,omitempty"`
}

// Create a new header with default values.
// This header does not report click events by default.
func NewHeader() *Header {
	return &Header{
		Version:     1,
		ClickEvents: false,
		ContSignal:  syscall.SIGCONT,
		StopSignal:  syscall.SIGSTOP,
	}
}
