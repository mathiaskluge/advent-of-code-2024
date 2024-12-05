package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	strings, err := readData("inputs/day04.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	wordMatrix := stringToMatrix(strings)

	wordMatrix45 := rotate45(wordMatrix)
	wordMatrix90 := rotate90(wordMatrix)
	wordMatrix135 := rotate45(wordMatrix90)

	var matchCount int

	matchCount += horizontalMatch(wordMatrix)
	fmt.Printf("Horizontal matches: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix90)
	fmt.Printf("With Vertical (rotate 90 degrees) matches: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix45)
	fmt.Printf("With Diagonal Positive matches: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix135)
	fmt.Printf("With Diagonal Negative matches: %d\n", matchCount)

}

func horizontalMatch(matrix [][]rune) int {
	var matchCount int

	for _, row := range matrix {
		for x := 0; x <= len(row)-4; x++ {
			// Forward Match
			if string(row[x:x+4]) == "XMAS" {
				matchCount++
				//fmt.Println(string(row[x : x+4]))
			}
			// Backward Match
			if x >= 4 && string(row[x-4:x]) == "SAMX" {
				matchCount++
				//fmt.Println(string(row[x-4 : x]))
			}
		}
	}
	return matchCount
}

func rotate45(matrix [][]rune) [][]rune {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]rune, rows+cols-1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			diagonalIndex := i + j
			result[diagonalIndex] = append([]rune{matrix[i][j]}, result[diagonalIndex]...)
		}
	}

	return result
}

func rotate90(matrix [][]rune) [][]rune {
	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}

func stringToMatrix(strings []string) [][]rune {
	matrix := make([][]rune, len(strings))

	for i, line := range strings {
		matrix[i] = []rune(line)
	}
	return matrix
}

func readData(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		line := scanner.Text()

		words = append(words, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
		return nil, err
	}
	return words, nil
}
