package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	data := "data"
	expected := "atad"

	actual := reverseString(data)

	assert.Equal(t, expected, actual)
}

func TestGetMode(t *testing.T) {
	expected := "dev"
	actual := ""

	if os.Getenv("MODE") == "dev" {
		actual = "dev"
	}
	assert.Equal(t, expected, actual)
}
