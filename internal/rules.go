package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// applyConversion берет последнее добавленное слово и переводит его в десятичную систему
func applyConversion(words []string, base int) {
	if len(words) == 0 {
		return
	}

	idx := len(words) - 1

	val, err := strconv.ParseInt(words[idx], base, 64)
	if err == nil {
		words[idx] = fmt.Sprintf("%d", val) // Без указателей всё выглядит просто
	}
}

// applyCase применяет функцию transform к N последним словам в массиве
func applyCase(words []string, n int, transform func(string) string) {
	length := len(words)
	if length == 0 || n <= 0 {
		return
	}

	if n > length {
		n = length
	}

	for i := length - n; i < length; i++ {
		words[i] = transform(words[i]) // Просто берем слово, меняем и кладем обратно
	}
}

// Capitalize делает первую букву заглавной
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	return strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
}

// FormatPunctuation приклеивает знаки препинания к предыдущему слову
func FormatPunctuation(text string) string {
	// Убираем пробел ПЕРЕД знаками препинания (.,!?:;)
	reBefore := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = reBefore.ReplaceAllString(text, "$1")

	// Убеждаемся, что ПОСЛЕ знака препинания есть пробел
	// (если это не конец строки и не другой знак препинания)
	reAfter := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	return text
}

// FormatQuotes форматирует одинарные кавычки (приклеивает их к словам внутри)
func FormatQuotes(text string) string {
	// Ищем текст внутри одинарных кавычек и убираем пробелы по краям внутри кавычек
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}

// FormatArticles заменяет 'a' и 'A' на 'an' и 'An', если дальше идет гласная или 'h'
func FormatArticles(text string) string {
	// Ищем 'a' или 'A', стоящие как отдельное слово, после которых идет гласная или 'h'
	re := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)

	// ReplaceAllStringFunc позволяет нам изменить найденный кусок текста
	return re.ReplaceAllStringFunc(text, func(match string) string {
		if strings.HasPrefix(match, "a") {
			// Берем "an " + последнюю букву из совпадения
			return "an " + match[len(match)-1:]
		}
		return "An " + match[len(match)-1:]
	})
}
