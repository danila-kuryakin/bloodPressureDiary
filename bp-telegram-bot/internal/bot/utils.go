package bot

import (
	"fmt"
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

func getIntFromBuffer(buf map[string]any, key string) (int, error) {
	val, exists := buf[key]
	if !exists {
		return 0, fmt.Errorf("key not found")
	}

	switch v := val.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("unsupported type %T", v)
	}
}
