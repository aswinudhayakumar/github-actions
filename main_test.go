package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStringSuccess(t *testing.T) {
	data := "data"
	expected := "ata"

	actual := reverseString(data)

	assert.Equal(t, expected, actual)
}
