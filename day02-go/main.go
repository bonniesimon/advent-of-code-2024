package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateReportSafety(reportLine string) {
	report := processReportStringIntoInts(reportLine)
	fmt.Println(report)
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

	// safeReports := 0

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() && i < 1 {
		i++
		fmt.Println(scanner.Text())
		evaluateReportSafety(scanner.Text())
	}
}
