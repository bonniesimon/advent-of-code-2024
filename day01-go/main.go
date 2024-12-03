package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func distance(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

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

	// var rightFreq map[int]int
	rightFreq := make(map[int]int)

	for _, row := range rows {
		// fmt.Println(strings.Split(row, "   "))
		col := strings.Split(row, "   ")
		leftNumber, _ := strconv.Atoi(col[0])
		rightNumber, _ := strconv.Atoi(col[1])

		left = append(left, leftNumber)

		value, ok := rightFreq[rightNumber]
		if !ok {
			rightFreq[rightNumber] = 1
		} else {
			rightFreq[rightNumber] = value + 1
		}
	}

	left = bubbleSort(left)

	// var distances []int

	similarityScore := 0

	for i := range left {
		val := rightFreq[left[i]]

		similarityScore = similarityScore + left[i]*val
	}

	// fmt.Println(distances)
	fmt.Println(similarityScore)
}
