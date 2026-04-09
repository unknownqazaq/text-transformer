package main

import (
	"fmt"
	"os"
	"text-transformer/internal"
)

func main() {
	// os.Args хранит аргументы командной строки
	inputFile, outputFile, err := validateArgs(os.Args)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1) // Завершаем программу с кодом ошибки
	}

	content, err := internal.ReadTextFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтение файла", err)
		os.Exit(1)
	}
	processedText := internal.Process(content)

	err = internal.WriteTextFile(outputFile, processedText)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		os.Exit(1)
	}

	fmt.Println("Шаг 3: Файл успешно прочитан и записан!")

	fmt.Println("Входной файл:", inputFile)
	fmt.Println("Выходной файл:", outputFile)
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
