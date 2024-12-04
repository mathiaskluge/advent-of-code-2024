package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read data by line
	data, err := readData("inputs/day03.txt")
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	// Join data into a single string
	dataString := strings.Join(data, "")

	// Part 1:
	// Find all mul(X,X) where X is a 1-3 digit number.
	// Multiply X*X and add to a total sum

	// find matches and put in format [["mul(123,456)", "123", "456"]]
	regex := `mul\((\d{1,3}),(\d{1,3})\)`
	r := regexp.MustCompile(regex)

	matches := r.FindAllStringSubmatch(dataString, -1)

	var sum int64

	for _, match := range matches {
		// fmt.Printf("%v * %v\n", match[1], match[2])

		product, err := mulStrings(match[1], match[2])
		if err != nil {
			fmt.Println("Error multiplying strings:", err)
			return
		}
		sum += int64(product)
		// fmt.Printf("%v \n", sum)
	}

	fmt.Printf("Part 1: %v\n", sum)

	// Part 2:
	// There are do() and don't() statements in the data.
	// Only compute the mul(X,X) after a do() block.
	// Skip computation in a don't() block.

	// adds 2 additional groups in front for do() and don't()
	regex = `(do\(\))|(don't\(\))|mul\((\d{1,3}),(\d{1,3})\)`
	r = regexp.MustCompile(regex)

	matches = r.FindAllStringSubmatch(dataString, -1)

	mulEnabled := true
	var sum2 int64

	for _, match := range matches {

		if mulEnabled && match[2] == "don't()" {
			mulEnabled = false
			// fmt.Printf("[2]: %v --> mulling disabled\n", match[2])
		}

		if mulEnabled && match[3] != "" && match[4] != "" {

			product, err := mulStrings(match[3], match[4])
			if err != nil {
				fmt.Println("Error multiplying strings:", err)
				return
			}

			sum2 += int64(product)

			// fmt.Printf("[3]: %v * [4]: %v ---> Current Sum: %v \n", match[3], match[4], sum2)

		}

		if !mulEnabled && match[1] == "do()" {
			mulEnabled = true
			// fmt.Printf("[1]: %v --> mulling enabled\n", match[1])
		}

	}

	fmt.Printf("Part 2: %v\n", sum2)

}

func mulStrings(a, b string) (int, error) {
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return aInt * bInt, nil
}

func readData(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
