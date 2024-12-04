package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	NONE Direction = "NONE"
	INC  Direction = "INC"
	DEC  Direction = "DEC"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func findDirection(difference int) Direction {
	if difference < 0 {
		return DEC
	}

	return INC
}

func evaluateReportSafety(reportLine string) bool {
	report := processReportStringIntoInts(reportLine)

	var difference int

	var direction Direction = NONE

	for i := 0; i < len(report)-1; i++ {
		difference = report[i+1] - report[i]
		if abs(difference) == 0 || abs(difference) > 3 {
			return false
		}

		if direction == NONE {
			direction = findDirection(difference)
		}

		if direction != findDirection(difference) {
			return false
		}
	}

	return true
}

func processReportStringIntoInts(reportLine string) []int {
	var reportInts []int

	reportStrings := strings.Split(reportLine, " ")

	for i := range reportStrings {
		reportInt, _ := strconv.Atoi(reportStrings[i])
		reportInts = append(reportInts, reportInt)
	}

	return reportInts
}

func main() {
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
	defer file.Close()

	safeReports := 0
	var isReportSafe bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		isReportSafe = evaluateReportSafety(scanner.Text())

		if isReportSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
