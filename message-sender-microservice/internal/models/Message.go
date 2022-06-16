package models

import(
	"encoding/json"
	"io"
)

type Message struct {
   Message string `json:"message"`
}

func (m *Message) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(m)
}