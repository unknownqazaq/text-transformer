package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	createFile, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Unable to creat file:", err)
		os.Exit(1)

	}

	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
		createFile.WriteString(string(data[:n]))
	}

	defer createFile.Close()

	fmt.Println("Done")

}
