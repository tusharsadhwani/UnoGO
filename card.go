package main

import "fmt"

type card struct {
	cardType string
	color    string
	value    int
}

func (c card) print() {
	switch c.cardType {
	case "Draw Four", "Wildcard":
		fmt.Printf("%s", c.cardType)
	case "Reverse", "Skip", "Draw Two":
		fmt.Printf("%s %s", c.color, c.cardType)
	case "Regular":
		fmt.Printf("%s %d", c.color, c.value)
	default:
		fmt.Printf("OOF %s %s %d", c.color, c.cardType, c.value)
	}
}
