package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInputFile() *os.File {
	if len(os.Args) < 2 {
		fmt.Println("Error: filename not given.\n\nUsage: ./your_program.sh [filename]")
		os.Exit(1)
	}
	inputFilePath := os.Args[1]
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	return file
}

func main() {
	file := parseInputFile()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
