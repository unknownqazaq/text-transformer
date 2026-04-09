package internal

import (
	"strings"
)

// Process принимает сырой текст, применяет правила и возвращает готовый текст
func Process(text string) string {
	// strings.Fields удобно разбивает текст на слова, игнорируя лишние пробелы и переносы
	words := strings.Fields(text)
	var processed []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "(hex)":
			applyConversion(&processed, 16)
		case "(bin)":
			applyConversion(&processed, 2)
		default:
			// Если это обычное слово, просто кладем его в итоговый массив
			processed = append(processed, word)
		}
	}

	// Склеиваем слова обратно через один пробел
	return strings.Join(processed, " ")
}
