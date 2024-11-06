package share

import (
	"fmt"
	"regexp"
)

func ExtractPRNumber(url string) (string, error) {
	re := regexp.MustCompile(`/pull/(\d+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", fmt.Errorf("PR number not found in URL")
	}
	return matches[1], nil
}
