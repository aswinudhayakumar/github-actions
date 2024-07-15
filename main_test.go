package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	
	data := "data"
	expected := "atad"

	actual := reverseString(data)

	assert.Equal(t, expected, actual)
}
