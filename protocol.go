package swaybarprotocol

import (
	"encoding/json"
	"log"
	"os"
	"errors"
)

// The swaybar protocol status
type Status struct {
	out      *os.File
	encoder  *json.Encoder
	header   *Header
	elements []Element
	Events   chan ClickEvent
}

func New() *Status {
	return &Status{
		out:     os.Stdout,
		encoder: json.NewEncoder(os.Stdout),
		header:  NewHeader(),
		Events: make(chan ClickEvent),
	}
}

// Initializes the status protocol by printing a header.
func (s *Status) Init() {
	err := s.encoder.Encode(s.header)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.WriteString("[")
	if err != nil {
		log.Fatal(err)
	}
}

// Add a body element
func (s *Status) Add(e *Element) {
	s.elements = append(s.elements, *e)
}

// Update and output each registered block
func (s *Status) Update() {
	err := s.encoder.Encode(s.elements)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.WriteString(",")
	if err != nil {
		log.Fatal(err)
	}
}

// Read click events
func (s *Status) ReadEvents() error {
	decoder := json.NewDecoder(os.Stdin)

	// Consume opening '['
	_, err := decoder.Token()
	if err != nil {
		return err
	}

	for decoder.More() {
		var event ClickEvent
		err = decoder.Decode(&event)
		if err != nil {
			return err
		}
		s.Events <- event
	}

	return errors.New("stdin exhausted")
}

func (s *Status) Write(p []byte) (n int, err error) {
	return s.out.Write(p)
}

func (stat *Status) WriteString(s string) (n int, err error) {
	return stat.Write([]byte(s))
}
