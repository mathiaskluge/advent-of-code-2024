package main

import (
	"testing"
)

func TestRotate45(t *testing.T) {
	input := [][]rune{
		{'a', 'b', 'c', 'd'},
		{'e', 'f', 'g', 'h'},
		{'i', 'j', 'k', 'l'},
	}
	expected := [][]rune{
		{'a'},
		{'e', 'b'},
		{'i', 'f', 'c'},
		{'j', 'g', 'd'},
		{'k', 'h'},
		{'l'},
	}

	result := rotate45(input)

	if !equalSlices(result, expected) {
		t.Errorf("rotate45Rune(%v) = %v; want %v", input, result, expected)
	}

}
func TestRotate90(t *testing.T) {
	input := [][]rune{
		{'a', 'b', 'c', 'd'},
		{'e', 'f', 'g', 'h'},
		{'i', 'j', 'k', 'l'},
	}
	expected := [][]rune{
		{'i', 'e', 'a'},
		{'j', 'f', 'b'},
		{'k', 'g', 'c'},
		{'l', 'h', 'd'},
	}
	result := rotate90(input)
	if !equalSlices(result, expected) {
		t.Errorf("rotate90Rune(%v) = %v; want %v", input, result, expected)
	}

}
func TestStringsToMatrix(t *testing.T) {
	input := []string{"abcd", "efgh", "ijkl", "mnop"}
	expected := [][]rune{
		{'a', 'b', 'c', 'd'},
		{'e', 'f', 'g', 'h'},
		{'i', 'j', 'k', 'l'},
		{'m', 'n', 'o', 'p'},
	}
	result := stringToMatrix(input)
	if !equalSlices(result, expected) {
		t.Errorf("stringToMatrix(%v) = %v; want %v", input, result, expected)
	}
}

func equalSlices(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalRunes(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
