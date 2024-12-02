package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reports, err := readReports("inputs/day02.txt")
	if err != nil {
		fmt.Println("error reading reports:", err)
		return
	}

	// Part 1:
	// Count all safe reports

	count := 0

	for i, report := range reports {
		isSafe := processReport(report)

		fmt.Printf("Report %d: %v -- Is Safe:%v\n", i, report, isSafe)
		if isSafe {
			count++
		}
	}
	fmt.Printf("%v/%d reports are safe\n", count, len(reports))
}

// Get report [1, 2, 3, 4, 5]
// 1. Check if is sorted
// --> Yes
// .   2. Check for duplicates
// .   .   Yes: return yes, check for only one and remove it
// .   .   No: analyze
// --> No
// .   2. check for duplicates
// .   .   if yes, direct false (cannot be dampened twice)
// .   .   if no, remove single unsorted element

func processReport(report []int) bool {

	isSorted := slices.IsSorted(report)
	duplicateCount := hasDuplicates(report)
	hasGapFirst, hasGapLast := hasLevelGap(report)

	// If more than one duplicate -> cant be corrected, return false
	if duplicateCount > 1 {
		return false
	}

	// If unsorted and has duplicates -> can't be corrected, return false
	if isSorted && duplicateCount == 0 {
		// Handle gaps at the first and last positions
		if hasGapFirst && hasGapLast {
			return false
		}

		if hasGapLast {
			report = report[:len(report)-1]
		} else if hasGapFirst {
			report = report[1:]
		}

		isSafe, _ := analyzeReport(report)
		return isSafe
	}

	// Sorted and 1 duplicate -> remove it,nalyze, return result
	if isSorted && duplicateCount == 1 {
		report = removeSingleDuplicate(report)
		isSafe, _ := analyzeReport(report)
		return isSafe
	}

	// Unsorted and no duplicate -> remove single unsorted element, analyze, return result
	if isSorted && duplicateCount == 0 {
		report = fixSingleUnsortedElement(report)
		isSafe, _ := analyzeReport(report)
		return isSafe
	}

	return true
}

func hasDuplicates(nums []int) int {

	duplicateCount := 0
	// Count occurances of each number
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
		if counts[num] > 1 {
			duplicateCount++
		}
	}
	return duplicateCount
}

func hasLevelGap(nums []int) (first bool, last bool) {
	if nums[0]-nums[1] > 3 || nums[0]-nums[1] < 3 {
		first = true
	}
	if nums[len(nums)-1]-nums[len(nums)-2] > 3 || nums[len(nums)-1]-nums[len(nums)-2] < 3 {
		last = true
	}
	return
}

func fixSingleUnsortedElement(nums []int) []int {
	// Helper function to check if the array is sorted
	isSortedInt := func(arr []int) bool {
		return sort.IntsAreSorted(arr)
	}

	// If the array is already sorted, return it as-is
	if isSortedInt(nums) {
		return nums
	}

	// Find the single out-of-place number
	outOfPlaceIndex := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			// If outOfPlaceIndex is already set, more than one issue exists
			if outOfPlaceIndex != -1 {
				return nums // Multiple issues; return unchanged
			}
			outOfPlaceIndex = i
		}
	}

	// If outOfPlaceIndex is valid, try removing the offending number
	if outOfPlaceIndex != -1 {
		// Try removing the out-of-place number
		candidate := append([]int{}, nums[:outOfPlaceIndex]...)
		candidate = append(candidate, nums[outOfPlaceIndex+1:]...)

		// Check if the new array is sorted
		if isSortedInt(candidate) {
			return candidate // Return the corrected array
		}

		// Handle special case: the out-of-place number might be earlier
		if outOfPlaceIndex > 0 {
			candidate = append([]int{}, nums[:outOfPlaceIndex-1]...)
			candidate = append(candidate, nums[outOfPlaceIndex:]...)
			if isSortedInt(candidate) {
				return candidate
			}
		}
	}

	// Return the original array if no single unsorted number could be identified
	return nums
}

func removeSingleDuplicate(nums []int) []int {
	// Count occurances of each number
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	// Check if there's only one duplicate
	duplicateCount := 0
	var duplicateNum int
	for num, count := range counts {
		if count > 2 {
			// More than one duplicate of the same number
			return nums
		}
		if count == 2 {
			duplicateCount++
			duplicateNum = num
		}
	}
	// More than one duplicated number
	if duplicateCount != 1 {
		return nums
	}

	// Remove duplicate
	result := []int{}
	duplicateRemoved := false
	for _, num := range nums {
		if num == duplicateNum && !duplicateRemoved {
			// Skip the first occurrence
			duplicateRemoved = true
		} else {
			result = append(result, num)
		}
	}

	return result
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

		// TODO: Address empty lines
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
