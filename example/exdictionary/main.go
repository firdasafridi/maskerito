package main

import (
	"encoding/json"
	"fmt"

	"github.com/firdasafridi/maskerito"
)

type Sample struct {
	Mobile []string `json:"mobile"`
	Bank   string
}

func main() {

	sample := Sample{
		Mobile: []string{"081313131313"},
		Bank:   "7123123123",
	}

	m, err := maskerito.New(maskerito.DefaultConfig())
	fmt.Println("is err?", err)

	c, err := m.Marshal(sample)

	output, _ := json.Marshal(c)
	fmt.Println(string(output), sample, err)
}
