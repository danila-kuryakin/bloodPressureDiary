package bot

import (
	"regexp"
)

func IsTwoOrThreeDigits(s string) bool {
	re := regexp.MustCompile(`^\d{2,3}$`)
	return re.MatchString(s)
}

func RemoveByID(keys []int, index int) []int {
	if index < 0 || index >= len(keys) {
		return keys
	}

	newKeys := append(keys[:index], keys[index+1:]...)
	return newKeys
}

//func RemoveByID[T any](slice []T, i int) []T {
//	// Проверки на корректность индекса
//	if i < 0 || i >= len(slice) {
//		return slice // или можно паниковать: panic("index out of range")
//	}
//
//	// Основной способ удаления
//	return append(slice[:i], slice[i+1:]...)
//}
