package bot

import (
	"strconv"
	"strings"
)

func splitInts(input string) []int {
	parts := strings.Split(input, ",")
	result := make([]int, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		if v, err := strconv.Atoi(p); err == nil {
			result = append(result, v)
		}
	}

	return result
}
