package normalizer

import "strings"

// Function to count the number of leading spaces in a string.
func countLeadingSpaces(s string) int {
	return len(s) - len(strings.TrimLeft(s, " "))
}

// Function to compute the greatest common divisor (GCD) of two numbers.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to determine the number of spaces used for indentation.
// It analyzes each line in the given text and computes the GCD
// for all non-zero indentations.
func detectIndent(lines []string) int {
	var indents []int

	for _, line := range lines {
		// Skip empty lines or lines that contain only spaces.
		if strings.TrimSpace(line) == "" {
			continue
		}
		spaces := countLeadingSpaces(line)
		// If the line has an indentation (i.e., starts with spaces), store the value.
		if spaces > 0 {
			indents = append(indents, spaces)
		}
	}

	// If no indentation is found, return 0.
	if len(indents) == 0 {
		return 0
	}

	// Compute the GCD for all indentation values.
	currentGCD := indents[0]
	for i := 1; i < len(indents); i++ {
		currentGCD = gcd(currentGCD, indents[i])
		// If the GCD becomes 1, further computation will not change the result.
		if currentGCD == 1 {
			break
		}
	}

	if currentGCD > 5 {
		for currentGCD > 5 {
			currentGCD /= 2
		}
	}

	return currentGCD
}
