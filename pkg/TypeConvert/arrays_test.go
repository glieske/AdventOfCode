package typeconvert

import (
	"reflect"
	"testing"
)

func TestArrayOfBytesToArrayOfStrings(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []string
	}{
		{
			input:    []byte("line1\nline2\nline3"),
			expected: []string{"line1", "line2", "line3"},
		},
		{
			input:    []byte("singleline"),
			expected: []string{"singleline"},
		},
		{
			input:    []byte(""),
			expected: []string{""},
		},
	}

	for _, test := range tests {
		result := ArrayOfBytesToArrayOfStrings(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ArrayOfBytesToArrayOfStrings(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
