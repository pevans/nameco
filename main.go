package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/moby/moby/pkg/namesgenerator"
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
		out = namesgenerator.GetRandomName(0)
	case "word":
		parts := strings.Split(namesgenerator.GetRandomName(0), "_")
		if len(parts) > 0 {
			out = parts[0]
		}
	}

	if out == "" {
		fmt.Println("there's no name to print... 🤔")
		os.Exit(1)
	}

	fmt.Println(out)
}
