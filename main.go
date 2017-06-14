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
	jsonData := &Message{"Hello, world!", false}
	b, _ := json.Marshal(jsonData)
	fmt.Println(string(b))
}
