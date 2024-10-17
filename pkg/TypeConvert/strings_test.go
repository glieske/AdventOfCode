package typeconvert

import (
	"testing"
)

func TestStringRuneToInt(t *testing.T) {
	tests := []struct {
		input    rune
		expected int
	}{
		{'0', 0},
		{'1', 1},
		{'9', 9},
	}

	for _, test := range tests {
		result := StringRuneToInt(test.input)
		if result != test.expected {
			t.Errorf("StringRuneToInt(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestStringRuneToString(t *testing.T) {
	tests := []struct {
		input    rune
		expected string
	}{
		{'a', "a"},
		{'b', "b"},
		{'1', "1"},
	}

	for _, test := range tests {
		result := StringRuneToString(test.input)
		if result != test.expected {
			t.Errorf("StringRuneToString(%q) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestStringRuneToInt32(t *testing.T) {
	tests := []struct {
		input    rune
		expected int32
	}{
		{'0', 0},
		{'1', 1},
		{'9', 9},
	}

	for _, test := range tests {
		result := StringRuneToInt32(test.input)
		if result != test.expected {
			t.Errorf("StringRuneToInt32(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestStringRuneToInt64(t *testing.T) {
	tests := []struct {
		input    rune
		expected int64
	}{
		{'0', 0},
		{'1', 1},
		{'9', 9},
	}

	for _, test := range tests {
		result := StringRuneToInt64(test.input)
		if result != test.expected {
			t.Errorf("StringRuneToInt64(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
