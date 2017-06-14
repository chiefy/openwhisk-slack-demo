package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	b, _ := json.Marshal(&Message{"Hello, world!"})
	fmt.Print(string(b))
}
