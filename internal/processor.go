package internal

import (
	"strconv"
	"strings"
)

// Process принимает сырой текст, применяет правила и возвращает готовый текст
func Process(text string) string {
	words := strings.Fields(text)
	var processed []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "(hex)":
			applyConversion(processed, 16)
		case "(bin)":
			applyConversion(processed, 2)
		case "(up)":
			applyCase(processed, 1, strings.ToUpper)
		case "(low)":
			applyCase(processed, 1, strings.ToLower)
		case "(cap)":
			applyCase(processed, 1, Capitalize)
		case "(up,", "(low,", "(cap,":
			// Проверяем следующее слово (там должно быть число с закрывающей скобкой)
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				nStr := strings.TrimSuffix(words[i+1], ")")
				n, err := strconv.Atoi(nStr)

				if err == nil {
					switch word {
					case "(up,":
						applyCase(processed, n, strings.ToUpper)
					case "(low,":
						applyCase(processed, n, strings.ToLower)
					case "(cap,":
						applyCase(processed, n, Capitalize)
					}

					i++ // Пропускаем число, так как мы его уже обработали
					continue
				}
			}
			processed = append(processed, word)
		default:
			// Обычные слова просто добавляем в список
			processed = append(processed, word)
		}
	}

	// Склеиваем слова обратно через один пробел
	result := strings.Join(processed, " ")

	// Применяем финальное форматирование пунктуации, кавычек и артиклей
	result = FormatPunctuation(result)
	result = FormatQuotes(result)
	result = FormatArticles(result)

	return result
}
