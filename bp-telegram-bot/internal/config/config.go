package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnv загружает .env файл в окружение. Строки .env файла:
// TG_TOKEN - Токен телеграмма
// BP_API_URL - URL сервиса с api (пример: localhost:8090)
func LoadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// пропускаем комментарии и пустые строки
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}
}
