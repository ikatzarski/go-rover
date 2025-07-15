package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	type testCase struct {
		name     string
		command  string
		expected string
	}

	testCases := []testCase{
		{name: "No command", command: "", expected: "0:0:N"},
		{name: "Rotate right", command: "R", expected: "0:0:E"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rover := NewRover()
			actual := rover.Execute(tc.command)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
