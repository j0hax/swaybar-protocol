package swaybarprotocol

import "encoding/json"

// Element wraps a status body.
// It contains metadata and helper functions.
type Element struct {
	body   *Body
	update func(b *Body)
}

// Create a new body
func NewElement(name string, updateFunc func(b *Body)) *Element {
	return &Element{
		body: &Body{
			Name: name,
		},
		update: updateFunc,
	}
}

// Updates the body and sends it as json
func (e *Element) MarshalJSON() ([]byte, error) {
	e.update(e.body)
	return json.Marshal(e.body)
}
