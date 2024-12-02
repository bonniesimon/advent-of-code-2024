package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: filename not given.\n\nUsage: ./your_program.sh [filename]")
		os.Exit(1)
	}
	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}

	rows := strings.Split(string(data), "\n")

	// Ignore empty array at the end due to EOF
	rows = rows[:len(rows)-1]

	var left []int
	var right []int

	for _, row := range rows {
		fmt.Println(strings.Split(row, "   "))
		col := strings.Split(row, "   ")
		leftNumber, _ := strconv.Atoi(col[0])
		rightNumber, _ := strconv.Atoi(col[1])

		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	fmt.Println(left)
	fmt.Println(right)
}
