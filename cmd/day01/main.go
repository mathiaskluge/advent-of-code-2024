package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "log"
    "sort"
)

func main() {

    // Read and parse file into 2 separate lists
    left, right, err := readFile("inputs/day01.txt")
    if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
    
    // Sort lists
	sort.Ints(left)
	sort.Ints(right)

    // Summarize list item differences
    sum, err := sumDiff(left, right)
    if err != nil {
        log.Fatalf("Error computing differences: %v", err)
    }

    fmt.Println(sum)
}

func sumDiff(left []int, right []int) (int64, error) {
    
    // Ensure lists are same length
    if len(left) != len(right) {
		return 0, fmt.Errorf("Lists are not the same length.")
	}

    var sum int64

    for i := range left {
        diff := left[i] - right[i]
        if diff < 0 {
            diff = -diff
        }
        sum += int64(diff)
    }
    return sum, nil
}


func readFile(filename string) ([]int, []int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    var  left, right []int
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) != 2 {
            return nil, nil, fmt.Errorf("Invalid format: %s", line)
        }

        id1, err := strconv.Atoi(parts[0])
        if err != nil {
            return nil, nil, err
        }

        id2, err := strconv.Atoi(parts[1])
        if err != nil {
            return nil, nil, err
        }

        left = append(left, id1)
        right = append(right, id2)
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return left, right, nil
}


