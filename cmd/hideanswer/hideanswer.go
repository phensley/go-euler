package main

import (
	"fmt"
	"os"

	"github.com/phensley/go-euler"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <problem> <answer>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(euler.HideAnswer(os.Args[1], os.Args[2]))
}
