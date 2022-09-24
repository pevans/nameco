package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMobyName(t *testing.T) {
	assert.NotEmpty(t, MobyName())
}

func TestMobyWord(t *testing.T) {
	assert.NotEmpty(t, MobyWord())
}
