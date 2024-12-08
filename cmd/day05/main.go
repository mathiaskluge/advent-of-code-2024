package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	instructions, updates, err := readData("inputs/day05.txt")
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	validUpdates := validateUpdates(instructions, updates)

	sum, err := sumMiddlePages(validUpdates)
	if err != nil {
		fmt.Println("error computing sum:", err)
		return
	}

	fmt.Println(sum)

}

func sumMiddlePages(updates [][]string) (int, error) {
	sum := 0

	for _, update := range updates {

		if len(update)%2 == 0 {
			panic("Update is not uneven")
		}

		pageNumber, err := strconv.Atoi(update[len(update)/2+1])
		if err != nil {
			fmt.Println("error converting to integer:", err)
			return 0, err
		}

		sum += pageNumber
	}

	return sum, nil
}

func validateUpdates(instructions map[string]*int, updates [][]string) [][]string {
	var validUpdates [][]string

	for _, update := range updates {
		for i := 0; i <= len(update)-2; i++ {
			if _, exists := instructions[update[i]+update[i+1]]; exists {
				validUpdates = append(validUpdates, update)
			}
		}
	}

	return validUpdates
}

func readData(filename string) (map[string]*int, [][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make(map[string]*int)
	var updates [][]string

	for scanner.Scan() {
		line := scanner.Text()

		// Skip blank lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		if len(line) == 5 {
			instructions[strings.ReplaceAll(line, "|", "")] = nil
		} else {
			page_numbers := strings.Split(line, ",")
			updates = append(updates, page_numbers)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
		return nil, nil, err
	}
	return instructions, updates, nil
}
