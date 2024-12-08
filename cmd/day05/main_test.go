package main

import (
	"os"
	"reflect"
	"testing"
)

func sumMiddlePagesTest(t *testing.T) {
	updates := [][]string{
		{"75", "47", "61", "53", "29"},
		{"97", "61", "53", "29", "13"},
		{"75", "29", "13"},
		{"75", "97", "47", "61", "53"},
		{"61", "13", "29"},
		{"97", "13", "75", "29", "47"},
	}

	expectedSum := 278

	result, err := sumMiddlePages(updates)
	if err != nil {
		t.Fatalf("Failed to build sum: %v", err)
	}

	if expectedSum != result {
		t.Errorf("Sum mismatch")
	}
}

func validateUpdatesTest(t *testing.T) {
	instructions := map[string]*int{
		"4753": nil,
		"9713": nil,
		"9761": nil,
		"9747": nil,
		"7529": nil,
		"6113": nil,
		"7553": nil,
		"2913": nil,
		"9729": nil,
		"5329": nil,
		"6153": nil,
		"9753": nil,
		"6129": nil,
		"4713": nil,
		"7547": nil,
		"9775": nil,
		"4761": nil,
		"7561": nil,
		"4729": nil,
		"7513": nil,
		"5313": nil,
	}

	updates := [][]string{
		{"75", "47", "61", "53", "29"},
		{"97", "61", "53", "29", "13"},
		{"75", "29", "13"},
		{"75", "97", "47", "61", "53"},
		{"61", "13", "29"},
		{"97", "13", "75", "29", "47"},
	}

	expectedUpdates := [][]string{
		{"75", "47", "61", "53", "29"},
		{"97", "61", "53", "29", "13"},
		{"75", "29", "13"},
	}

	result := validateUpdates(instructions, updates)

	if !reflect.DeepEqual(result, expectedUpdates) {
		t.Errorf("Valid Updates mismatch.\nExpected: %v\nGot: %v", expectedUpdates, updates)
	}

}

func TestReadData(t *testing.T) {
	testFileName := "test_data.txt"
	fileContent := `17|23
94|96
94|58
22|53

49,85,73,74,96,32,76,58,95,57,13,93,14,99,56,47,75
74,68,64,75,78,18,41,67,15
83,95,85,32,58,14,76,77,74,68,47,96,49,75,88,56,29,78,18`
	err := os.WriteFile(testFileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFileName)

	expectedInstructions := map[string]*int{
		"1723": nil,
		"9496": nil,
		"9458": nil,
		"2253": nil,
	}
	expectedUpdates := [][]string{
		{"49", "85", "73", "74", "96", "32", "76", "58", "95", "57", "13", "93", "14", "99", "56", "47", "75"},
		{"74", "68", "64", "75", "78", "18", "41", "67", "15"},
		{"83", "95", "85", "32", "58", "14", "76", "77", "74", "68", "47", "96", "49", "75", "88", "56", "29", "78", "18"},
	}

	instructions, updates, err := readData(testFileName)
	if err != nil {
		t.Fatalf("Error reading data: %v", err)
	}

	if !reflect.DeepEqual(instructions, expectedInstructions) {
		t.Errorf("Instructions mismatch.\nExpected: %v\nGot: %v", expectedInstructions, instructions)
	}
	if !reflect.DeepEqual(updates, expectedUpdates) {
		t.Errorf("Updates mismatch.\nExpected: %v\nGot: %v", expectedUpdates, updates)
	}
}
