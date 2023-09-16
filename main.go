package main

import (
	"fmt"
	"strings"
)

func main() {
	var command = "walk outside"
	var exit = strings.Contains(command, "wal")
	fmt.Println("you leave the cave:", exit)
}
