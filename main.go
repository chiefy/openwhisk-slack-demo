package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Text  string `json:"text"`
	Error bool   `json:"error"`
}

func main() {
	b, _ := json.Marshal(&Message{"Hello, world!", false})
	fmt.Print(string(b))
}
