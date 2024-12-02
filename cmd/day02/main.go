package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readReports("inputs/day02.txt")
	if err != nil {
		fmt.Println("error reading reports:", err)
		return
	}

	count := 0

	for i, report := range reports {
		isSafe, _ := analyzeReport(report)
		fmt.Printf("Report %d: %v -- Is Safe:%v\n", i, report, isSafe)
		if isSafe {
			count++
		}
	}
	fmt.Printf("%v/%d reports are safe\n", count, len(reports))

}

func analyzeReport(report []int) (bool, error) {
	// handle length < 2
	if len(report) < 2 {
		return true, nil
	}

	// Check for direction (increasing or decreasing)
	firstDiff := report[1] - report[0]

	// Handle initial zero
	if firstDiff == 0 {
		return false, nil
	}

	isIncreasing := firstDiff > 0

	// Check all pairs maintain the same direction
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// Check direction consistency
		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false, nil
		}

		// Check difference magnitude
		absDiff := diff
		if !isIncreasing {
			absDiff = -diff
		}
		if absDiff < 1 || absDiff > 3 {
			return false, nil
		}
	}

	return true, nil
}

func readReports(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataRow := []int{}

		report := scanner.Text()
		levels := strings.Fields(report)

		// Note: Add empty line hanlding
		for i := range levels {

			level, err := strconv.Atoi(levels[i])
			if err != nil {
				return nil, err
			}
			dataRow = append(dataRow, level)
		}
		data = append(data, dataRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
