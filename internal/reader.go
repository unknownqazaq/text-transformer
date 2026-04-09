package internal

import (
	"os"
)

// ReadTextFile читает всё содержимое файла и возвращает его в виде строки.
func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
