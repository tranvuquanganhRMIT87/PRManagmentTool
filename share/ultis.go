package share

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractPRNumber(url string) (string, error) {
	re := regexp.MustCompile(`/pull/(\d+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", fmt.Errorf("PR number not found in URL")
	}
	return matches[1], nil
}
func ExtractFirstChangedLine(patch string) int {
	for _, line := range strings.Split(patch, "\n") {
		if strings.HasPrefix(line, "@@") {
			return ParseLineNumber(line)
		}
	}
	return 0
}
func ParseLineNumber(line string) int {
	// Example parsing logic for extracting line number from patch line
	return 1
}
func Encode(text string) []int {
	return []int{1, 2, 3} // placeholder implementation
}
