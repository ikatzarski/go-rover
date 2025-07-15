package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	rover := NewRover()
	actual := rover.Execute("")
	expected := "0:0:N"

	assert.Equal(t, expected, actual)
}
