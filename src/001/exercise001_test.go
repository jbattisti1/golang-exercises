package main

import (
	"testing"
)

func TestIsDevisableBySeven(t *testing.T) {
	inputs := []struct {
		number   int
		expected bool
	}{
		{number: 0, expected: true},
		{number: 1, expected: false},
		{number: 2, expected: false},
		{number: 7, expected: true},
		{number: 2000, expected: false},
		{number: 2009, expected: true},
	}

	for _, input := range inputs {
		actual := IsDevisableBySeven(input.number)
		if input.expected != actual {
			t.Errorf("\"DevisableBySeven('%v')\" failed, expected -> %v, got -> %v", input.number, input.expected, actual)
		}
	}
}
