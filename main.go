package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text-transformer/internal"
	"time"
)

func main() {
	start := time.Now()

	inputFile, outputFile, err := validateArgs(os.Args)
	if err != nil {
		log.Println("Ошибка:", err)
		os.Exit(1)
	}

	content, err := internal.ReadTextFile(inputFile)
	if err != nil {
		log.Println("Ошибка при чтение файла", err)
		os.Exit(1)
	}
	log.Printf("Файл %s прочитан: символов=%d слов=%d", inputFile, len(content), len(strings.Fields(content)))

	processedText := internal.Process(content)
	log.Printf("Преобразования завершены: символов=%d слов=%d", len(processedText), len(strings.Fields(processedText)))

	err = internal.WriteTextFile(outputFile, processedText)
	if err != nil {
		log.Println("Ошибка при записи файла:", err)
		os.Exit(1)
	}

	log.Println("Входной файл:", inputFile)
	log.Println("Выходной файл:", outputFile)
	log.Printf("Обработка завершена успешно за %s", time.Since(start).Round(time.Millisecond))
}

// validateArgs проверяет правильность переданных аргументов
func validateArgs(args []string) (string, string, error) {
	// args[0] - это имя самой программы, args[1] и args[2] - файлы
	if len(args) != 3 {
		return "", "", fmt.Errorf("использование: go run . input.txt output.txt")
	}

	inputFile := args[1]
	outputFile := args[2]

	// Проверяем, существует ли входной файл
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return "", "", fmt.Errorf("входной файл не существует: %s", inputFile)
	}

	if outputFile == "" {
		return "", "", fmt.Errorf("имя выходного файла не может быть пустым")
	}

	return inputFile, outputFile, nil
}
