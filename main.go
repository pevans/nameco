package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := "name"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	var out string

	switch command {
	case "name":
		out = MobyName()
	case "word":
		out = MobyWord()
	}

	if out == "" {
		fmt.Println("there's no name to print... 🤔")
		os.Exit(1)
	}

	fmt.Println(out)
}
