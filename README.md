# maskerito

`maskerito` is masking library for golang. Especially for indonesia dictionary. 

Library maskerito provides a library to do masking struct and dictionary with go filed. 

## Overview
[![Go Doc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/firdasafridi/maskerito)
[![Go Report Card](https://goreportcard.com/badge/github.com/firdasafridi/maskerito)](https://goreportcard.com/report/github.com/firdasafridi/maskerito)


## The Idea
### Sample JSON Payload
Before the masking:
```
{
    "mobile": [
        "081313131313"
    ],
    "Bank": "7123123423"
}
```
After the encryption:
```
{
    "mobile": [
        "0813****1313"
    ],
    "Bank": "7123****23"
}
```
### Struct Tag Field Implementation
```go
type A struct {
    A string    `masker:"mobile"`    
}
```

## Full Sample
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/firdasafridi/maskerito"
)

type Sample struct {
	Mobile []string `masker:"mobile"`
	Bank   string   `masker:"bank"`
}

func main() {

	sample := Sample{
		Mobile: []string{"081313131313"},
		Bank:   "7123123123",
	}

	m, err := maskerito.New(maskerito.DefaultConfig())
	fmt.Println("is err?", err)

	c, err := m.Struct(sample)

	output, _ := json.Marshal(c)
	fmt.Println(string(output), sample, err)
}

```

## Limitation
`gocrypt` supports the dictionary type. Need more research & development to support the library for more dictionary type data.
