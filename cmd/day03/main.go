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

	// find matches and put in format [["mul(123,456)", "123", "456"]]
	regex := `mul\((\d{1,3}),(\d{1,3})\)`
	r := regexp.MustCompile(regex)

	matches := r.FindAllStringSubmatch(dataString, -1)

	var sum int64

	for _, match := range matches {
		// fmt.Printf("%v * %v\n", match[1], match[2])

		factor1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting factor1:", err)
			return
		}
		factor2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting factor2:", err)
			return
		}
		sum += int64(factor1) * int64(factor2)
		// fmt.Printf("%v \n", sum)
	}

	fmt.Printf("Final sum: %v\n", sum)
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
