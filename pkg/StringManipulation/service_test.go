package stringmanipulation

import (
	"errors"
	"testing"
)

func TestIsStringRuneANumber(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'0', true},
		{'5', true},
		{'9', true},
		{'a', false},
		{'-', false},
		{' ', false},
	}

	for _, test := range tests {
		result := IsStringRuneANumber(test.input)
		if result != test.expected {
			t.Errorf("IsStringRuneANumber(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestGetFirstNumberFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		err      error
	}{
		{"abc123", 1, nil},
		{"no numbers", -1, errors.New("no number found in string")},
		{"123abc", 1, nil},
		{"", -1, errors.New("no number found in string")},
	}

	for _, test := range tests {
		result, err := GetFirstNumberFromString(test.input)
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf(
				"GetFirstNumberFromString(%q) = %d, %v; want %d, %v",
				test.input,
				result,
				err,
				test.expected,
				test.err,
			)
		}
	}
}

func TestGetLastNumberFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		err      error
	}{
		{"abc123", 3, nil},
		{"no numbers", -1, errors.New("no number found in string")},
		{"123abc", 3, nil},
		{"", -1, errors.New("no number found in string")},
	}

	for _, test := range tests {
		result, err := GetLastNumberFromString(test.input)
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf(
				"GetLastNumberFromString(%q) = %d, %v; want %d, %v",
				test.input,
				result,
				err,
				test.expected,
				test.err,
			)
		}
	}
}
