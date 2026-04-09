package internal

import (
	"os"
)

// WriteTextFile записывает строку в файл. Если файла нет - он будет создан.
func WriteTextFile(filename string, content string) error {
	// 0644 - стандартные права доступа (чтение/запись для владельца, чтение для остальных)
	return os.WriteFile(filename, []byte(content), 0644)
}
