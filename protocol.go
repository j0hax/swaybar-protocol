package swaybarprotocol

import (
	"encoding/json"
	"io"
	"os"
)

// The swaybar protocol status
type Status struct {
	writer  io.Writer
	encoder *json.Encoder
	header  *Header
}

func New() *Status {
	return &Status{
		writer:  os.Stdout,
		encoder: json.NewEncoder(os.Stdout),
		header:  NewHeader(),
	}
}

// Initializes the status protocol by printing a header.
func (s *Status) Init() error {
	err := s.encoder.Encode(s.header)
	if err != nil {
		return err
	}
	_, err = s.WriteString("[")
	return err
}

func (s *Status) Output(items []Body) error {
	err := s.encoder.Encode(items)
	if err != nil {
		return err
	}
	_, err = s.WriteString(",")
	return err
}

/* TODO: implement in object-oriented fashion
func Read(r io.Reader) (*ClickEvent, error) {
	dec := json.NewDecoder(r)
	event := &ClickEvent{}
	err := dec.Decode(event)
	return event, err
}
*/

func (s *Status) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

func (stat *Status) WriteString(s string) (n int, err error) {
	return stat.Write([]byte(s))
}
