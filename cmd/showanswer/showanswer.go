package main

import (
	"fmt"
	"os"

	"github.com/phensley/go-euler"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <problem> <crypted>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(euler.RevealAnswer(os.Args[1], os.Args[2]))
}
