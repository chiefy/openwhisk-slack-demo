package main

import (
	"fmt"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	fmt.Print("{ \"msg\":\"hi\" }")
}
