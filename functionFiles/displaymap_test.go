package asciiart

import (
	"testing"
)

/*
TestDisplayAsciiArt verifies the DisplayAsciiArt function's ability to correctly assemble ASCII art from a given character map and input string.
This test checks if the function properly formats multiple characters into a single ASCII art representation, ensuring correct spacing and alignment.
*/
func TestDisplayAsciiArt(t *testing.T) {
	testMap := map[rune][]string{
		'A': {"1", "2", "3", "4", "5", "6", "7", "8"},
		'B': {"9", "10", "11", "13", "14", "15", "16", "17"},
	}
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test two chracters",
			input:    "AB",
			expected: "1 9 \n2 10 \n3 11 \n4 13 \n5 14 \n6 15 \n7 16 \n8 17 \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DisplayAsciiArt(testMap, tt.input)
			if result != tt.expected {
				t.Errorf("Test %s failed. Expected %q, got %q", tt.name, tt.expected, result)
			}
		})
	}
}
