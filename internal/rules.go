package internal

import (
	"fmt"
	"strconv"
)

// applyConversion берет последнее добавленное слово и переводит его в десятичную систему
func applyConversion(words *[]string, base int) {
	if len(*words) == 0 {
		return // Защита: если перед тегом нет никаких слов, ничего не делаем
	}

	idx := len(*words) - 1 // Получаем индекс предыдущего слова

	// Пытаемся превратить строку в число (base - это 16 для hex или 2 для bin)
	val, err := strconv.ParseInt((*words)[idx], base, 64)
	if err == nil {
		// Если конвертация прошла без ошибок, заменяем слово на новое значение
		(*words)[idx] = fmt.Sprintf("%d", val)
	}
}
