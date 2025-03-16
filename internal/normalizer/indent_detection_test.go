package normalizer

import (
	"testing"
)

func TestDetectIndent(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "Empty input",
			input:    []string{},
			expected: 0,
		},
		{
			name:     "Only blank lines",
			input:    []string{"", "   ", "\t"},
			expected: 0,
		},
		{
			name:     "No indentation",
			input:    []string{"file", "folder"},
			expected: 0,
		},
		{
			name:     "Uniform indent of 4 spaces",
			input:    []string{"    file1", "    file2"},
			expected: 4,
		},
		{
			name:     "Uniform indent of 4 spaces",
			input:    []string{"     file1", "     file2"},
			expected: 5,
		},
		{
			name:     "Check max tab size(6)",
			input:    []string{"      six spaces file"},
			expected: 3,
		},
		{
			name:     "Check max tab size(8)",
			input:    []string{"        six spaces file"},
			expected: 4,
		},
		{
			name:     "Mixed indent multiples (2, 4, 6 spaces)",
			input:    []string{"  folder", "    file", "      subfile"},
			expected: 2,
		},
		{
			name:     "Uniform indent of 3 spaces",
			input:    []string{"   item", "      subitem"},
			expected: 3,
		},
		{
			name:     "Mixed with tab lines (tabs are ignored)",
			input:    []string{"\tfile", "    folder"},
			expected: 4,
		},
		{
			name:     "Mixed with spaces and empty lines",
			input:    []string{"  file", "", "   ", "    folder", "folder"},
			expected: 2,
		},
		{
			name:     "Mixed with spaces and mixed indent (spaces then tab)",
			input:    []string{"  \tfile", "    file"},
			expected: 2,
		},
		{
			name:     "Only tabs",
			input:    []string{"\tfile", "\t\tfolder"},
			expected: 0,
		},
		{
			name:     "Non-multiples indent (3 and 2 spaces, GCD = 1)",
			input:    []string{"   file", "  folder"},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectIndent(tt.input)
			if got != tt.expected {
				t.Errorf("detectIndent(%v) = %d; expected %d", tt.input, got, tt.expected)
			}
		})
	}
}
