package normalizer

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Only empty lines",
			input:    "\n\n\n",
			expected: []string{},
		},
		{
			name:     "Lines with spaces",
			input:    "  \n   \n",
			expected: []string{},
		},
		{
			name:     "Decorative symbols",
			input:    "root/\n├── dir/\n│   └── file.txt",
			expected: []string{"root/", "\tdir/", "\t\tfile.txt"},
		},
		{
			name:     "|--",
			input:    "root/\n|-- dir/\n│   |-- file.txt",
			expected: []string{"root/", "\tdir/", "\t\tfile.txt"},
		},
		{
			name:     "Comments",
			input:    "dir/  # comment\nfile.txt",
			expected: []string{"dir/", "file.txt"},
		},
		{
			name:     "Tabs and spaces",
			input:    "\tdir/\n        file.txt",
			expected: []string{"dir/", "\tfile.txt"},
		},
		{
			name:     "Mixed case",
			input:    "├── dir/  # comment\n│   └── file.txt\n\n",
			expected: []string{"dir/", "\tfile.txt"},
		},
		{
			name:     "Multiline input",
			input:    "dir/\n  subdir/\n    file.txt",
			expected: []string{"dir/", "\tsubdir/", "\t\tfile.txt"},
		},
		{
			name:     "Unicode symbols",
			input:    "директория/\n  файл.txt",
			expected: []string{"директория/", "\tфайл.txt"},
		},
		{
			name:     "Trailing newline",
			input:    "dir/\nfile.txt\n",
			expected: []string{"dir/", "file.txt"},
		},
		{
			name:     "Root with decorative ",
			input:    "├── dir/\n│   └── file.txt",
			expected: []string{"dir/", "\tfile.txt"},
		},
		{
			name:     "Directory without slash",
			input:    "root\n|-- file",
			expected: []string{"root/", "\tfile"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Normalize(tt.input)
			if !equalSlices(result, tt.expected) {
				t.Errorf("Normalize() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
