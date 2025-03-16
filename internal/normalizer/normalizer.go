package normalizer

import (
	"regexp"
	"strings"
)

var (
	decorRegex   = regexp.MustCompile("^(?:[├└ │─|`*+\\]\\}\\)>-]+)")
	commentRegex = regexp.MustCompile(`\s*#.*`)
	emptyRegex   = regexp.MustCompile(`^[\s.]+$`)
)

func Normalize(input string) []string {
	lines := strings.Split(input, "\n")
	var result []string

	// first remove all decorative symbols and comments
	for i := range lines {
		// decorative symbols
		lines[i] = cleanLine(lines[i])

		// comments
		lines[i] = commentRegex.ReplaceAllString(lines[i], "")

		// remove all \r
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")
	}

	// find how many spaces go with a tab.
	spacesTabCount := detectIndent(lines)

	// replace spaces to tabs
	for _, line := range lines {
		newLine := line
		if spacesTabCount > 0 {
			newLine = strings.ReplaceAll(line, strings.Repeat(" ", spacesTabCount), "\t")
		}
		// remove empty or placeholder lines
		if newLine == "" || emptyRegex.MatchString(newLine) {
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

	// add slashes to directories
	for i := range result {
		if i+1 >= len(result) {
			break
		}
		if strings.HasSuffix(result[i], "/") {
			continue
		}

		if countLeadingTabs(result[i]) < countLeadingTabs(result[i+1]) {
			result[i] += "/"
		}
	}

	return result
}

// Function to count the number of leading spaces in a string.
func countLeadingTabs(s string) int {
	return len(s) - len(strings.TrimLeft(s, "\t"))
}

func cleanLine(line string) string {
	// remove all decorative symbols
	return decorRegex.ReplaceAllStringFunc(line, func(s string) string {
		// replace all decorative symbols with spaces
		return strings.Repeat(" ", len([]rune(s)))
	})
}
