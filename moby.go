package main

import (
	"strings"

	"github.com/moby/moby/pkg/namesgenerator"
)

func MobyName() string {
	return namesgenerator.GetRandomName(0)
}

func MobyWord() string {
	parts := strings.Split(MobyName(), "_")
	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}
