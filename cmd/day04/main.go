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
	fmt.Printf("0 Degrees: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix90)
	fmt.Printf("90 Degrees: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix45)
	fmt.Printf("45 Degrees: %d\n", matchCount)
	matchCount += horizontalMatch(wordMatrix135)
	fmt.Printf("135 Degrees: %d\n", matchCount)

	fmt.Printf("\n Matching X-MAS patterns\n")

	matchCount = patternMatch(wordMatrix)
	fmt.Println(matchCount)
}

func patternMatch(matrix [][]rune) int {
	var matchCount int

	// limits at [:-1][:-1] since X is build from the center letter
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {

			// every potentioal match has a center 'A'
			if matrix[i][j] != 'A' {
				continue
			}

			leftLeg := string([]rune{
				matrix[i-1][j-1], // X . .
				matrix[i][j],     // . j .
				matrix[i+1][j+1], // . . X
			})
			rightLeg := string([]rune{
				matrix[i-1][j+1], // . . X
				matrix[i][j],     // . j .
				matrix[i+1][j-1], // X . .
			})

			if (leftLeg == "MAS" || leftLeg == "SAM") && (rightLeg == "MAS" || rightLeg == "SAM") {
				matchCount++
			}
		}
	}

	return matchCount
}

func horizontalMatch(matrix [][]rune) int {
	var matchCount int

	for _, row := range matrix {
		if len(row) < 4 {
			continue
		}
		for x := 0; x <= len(row)-4; x++ {
			substring := string(row[x : x+4])
			if substring == "XMAS" || substring == "SAMX" {
				matchCount++
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
