package main

import (
	"testing"
)

func TestProblemDamper(t *testing.T) {
	// Test cases
	tests := []struct {
		report   []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},  // Safe without removing any level
		{[]int{1, 2, 7, 8, 9}, false}, // Unsafe regardless of which level is removed
		{[]int{9, 7, 6, 2, 1}, false}, // Unsafe regardless of which level is removed
		{[]int{1, 3, 2, 4, 5}, true},  // Safe by removing the second level (3)
		{[]int{8, 6, 4, 4, 1}, true},  // Safe by removing the third level (4)
		{[]int{1, 3, 6, 7, 9}, true},  // Safe without removing any level
		{[]int{3, 4, 5, 5, 5}, false}, // Unsafe due to too many duplicates
		{[]int{6, 4, 1, 1, 1}, false}, // Unsafe due to too many duplicates
	}

	for _, test := range tests {
		result := problemDamper(test.report)
		if result != test.expected {
			t.Errorf("problemDamper(%v) = %v; want %v", test.report, result, test.expected)
		}
	}
}

func TestAnalyzeReport(t *testing.T) {
	// Test cases
	tests := []struct {
		report   []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},  // Safe: all decreasing by 1 or 2
		{[]int{1, 2, 7, 8, 9}, false}, // Unsafe: increase of 5 between 2 and 7
		{[]int{9, 7, 6, 2, 1}, false}, // Unsafe: decrease of 4 between 6 and 2
		{[]int{1, 3, 2, 4, 5}, false}, // Unsafe: 1->3 increasing but 3->2 decreasing
		{[]int{8, 6, 4, 4, 1}, false}, // Unsafe: 4->4 neither increase nor decrease
		{[]int{1, 3, 6, 7, 9}, true},  // Safe: all increasing by 1, 2, or 3
	}

	for _, test := range tests {
		result, _ := analyzeReport(test.report)
		if result != test.expected {
			t.Errorf("analyzeReport(%v) = %v; want %v", test.report, result, test.expected)
		}
	}
}
