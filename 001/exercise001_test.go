package exercise001

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
		actual := isDevisableBySeven(input.number)
		if input.expected != actual {
			t.Errorf("\"isDevisableBySeven(%v)\" failed, expected -> %v, got -> %v", input.number, input.expected, actual)
		}
	}
}

func TestIsNotMutipleOfFive(t *testing.T) {
	inputs := []struct {
		number   int
		expected bool
	}{
		{number: 0, expected: false},
		{number: 1, expected: true},
		{number: 5, expected: false},
		{number: 7, expected: true},
		{number: 10, expected: false},
		{number: 100, expected: false},
		{number: 275, expected: false},
		{number: 278, expected: true},
		{number: 2000, expected: false},
	}

	for _, input := range inputs {
		actual := isNotMultipleOfFive(input.number)
		if input.expected != actual {
			t.Errorf("\"isNotMultipleOfFive(%v)\" failed, expected -> %v, got -> %v", input.number, input.expected, actual)
		}
	}
}
