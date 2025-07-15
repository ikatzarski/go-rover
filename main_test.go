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
		{name: "Rotate right twice", command: "RR", expected: "0:0:S"},
		{name: "Rotate right three times", command: "RRR", expected: "0:0:W"},
		{name: "Rotate right four times", command: "RRRR", expected: "0:0:N"},
		{name: "Rotate left", command: "L", expected: "0:0:W"},
		{name: "Rotate left twice", command: "LL", expected: "0:0:S"},
		{name: "Rotate left three times", command: "LLL", expected: "0:0:E"},
		{name: "Rotate left four times", command: "LLLL", expected: "0:0:N"},
		{name: "Rotate right and left", command: "RL", expected: "0:0:N"},
		{name: "Rotate right three times and left once", command: "RRRL", expected: "0:0:S"},
		{name: "Move forward North once", command: "M", expected: "0:1:N"},
		{name: "Move forward North three times", command: "MMM", expected: "0:3:N"},
		{name: "Move forward North and wrap around", command: "MMMMMMMMMM", expected: "0:0:N"},
		{name: "Move forward North, wrap around and move again", command: "MMMMMMMMMMM", expected: "0:1:N"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rover := NewRover()
			actual := rover.Execute(tc.command)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
