package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStringSuccess(t *testing.T) {
	data := "data"
	expected := "atad"

	actual := reverseString(data)

	assert.Equal(t, expected, actual)
}
