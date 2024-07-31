package isogram

import (
	"fmt"
	"strings"
)

func IsIsogram(s string) bool {
	charMap := make(map[rune]bool)

	for _, char := range strings.ToLower(s) {
		if char == ' ' || char == '-' {
			continue
		}

		if charMap[char] {
			return false
		}

		charMap[char] = true
	}

	return true
}

func main() {
	// Test cases
	testCases := []string{
		"lumberjacks",
		"background",
		"downstream",
		"six-year-old",
		"isograms",
	}

	for _, testCase := range testCases {
		if IsIsogram(testCase) {
			fmt.Printf("%s is an isogram\n", testCase)
		} else {
			fmt.Printf("%s is not an isogram\n", testCase)
		}
	}
}
