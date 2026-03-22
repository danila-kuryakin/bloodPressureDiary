package bot

import (
	"regexp"
)

func IsTwoOrThreeDigits(s string) bool {
	re := regexp.MustCompile(`^\d{2,3}$`)
	return re.MatchString(s)
}
