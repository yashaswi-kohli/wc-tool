package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var flag, filePath string
	if len(os.Args) < 3 {
		filePath = os.Args[1]
	} else {
		flag, filePath = os.Args[1], os.Args[2]
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	switch flag {
	case "-c":
		fmt.Println(numOfBytes(file), file.Name())
	case "-l":
		fmt.Println(numOfLines(file), file.Name())
	case "-w":
		fmt.Println(numOfWords(file), file.Name())
	case "-m":
		fmt.Println(numOfCharacters(file), file.Name())
	default:
		//* Reset file pointer before each function call
		file.Seek(0, 0)
		lines := numOfLines(file)

		file.Seek(0, 0)
		words := numOfWords(file)

		file.Seek(0, 0)
		bytes := numOfBytes(file)

		fmt.Println(lines, words, bytes, file.Name())
	}
}

func numOfBytes(file *os.File) int {
	totalBytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	return len(totalBytes)
}

func numOfLines(file *os.File) int {
	count, line := 0, bufio.NewScanner(file)

	for line.Scan() {
		count++
	}

	if line.Err() != nil {
		log.Fatal(line.Err())
	}

	return count
}

func numOfWords(file *os.File) int {
	count, input := 0, bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		words := strings.Fields(line)
		count += len(words)
	}

	if input.Err() != nil {
		log.Fatal(input.Err())
	}
	return count
}

func numOfCharacters(file *os.File) int {
	content, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	return utf8.RuneCount(content)
}
