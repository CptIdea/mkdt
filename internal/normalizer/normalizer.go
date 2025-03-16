package normalizer

import (
	"regexp"
	"strings"
)

var (
	decorRegex   = regexp.MustCompile(`[├│└──]`)
	commentRegex = regexp.MustCompile(`\s*#.*`)
)

func Normalize(input string) []string {
	lines := strings.Split(input, "\n")
	var result []string

	// first remove all decorative symbols and comments
	for i := range lines {
		// decorative symbols
		lines[i] = decorRegex.ReplaceAllString(lines[i], " ")

		// comments
		lines[i] = commentRegex.ReplaceAllString(lines[i], "")
	}

	// find how many spaces go with a tab.
	spacesTabCount := detectIndent(lines)

	// replace spaces to tabs
	for _, line := range lines {
		newLine := line
		if spacesTabCount > 0 {
			newLine = strings.ReplaceAll(line, strings.Repeat(" ", spacesTabCount), "\t")
		}
		if newLine == "" || strings.ReplaceAll(newLine, "\t", "") == "" || strings.ReplaceAll(newLine, " ", "") == "" {
			continue
		}
		result = append(result, newLine)
	}

	// if in first line is a tab, remove it
	if len(result) > 0 && strings.HasPrefix(result[0], "\t") {
		extraTabs := countLeadingTabs(result[0])
		for i := range result {
			result[i] = strings.Replace(result[i], "\t", "", extraTabs)
		}
	}

	return result
}

// Function to count the number of leading spaces in a string.
func countLeadingTabs(s string) int {
	return len(s) - len(strings.TrimLeft(s, "\t"))
}
